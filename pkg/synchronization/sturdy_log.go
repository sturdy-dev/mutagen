package synchronization

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/mutagen-io/mutagen/pkg/logging"
	"github.com/mutagen-io/mutagen/pkg/sturdy/api"
	sturdycontext "github.com/mutagen-io/mutagen/pkg/sturdy/context"
	"github.com/mutagen-io/mutagen/pkg/sturdy/debounce"
	"github.com/mutagen-io/mutagen/pkg/synchronization/core"
)

var (
	debounces   map[string]func(f func())
	debouncesMx sync.Mutex
	states      chan *State
)

// SturdyVersion is set at compile time
var SturdyVersion = "development"

func init() {
	debounces = make(map[string]func(f func()))
	states = make(chan *State, 1024)
	go stateReporter()
}

type stateTimestamp struct {
	Labels    map[string]string
	Status    Status
	Timestamp time.Time
}

var (
	statesMap   = map[string]*stateTimestamp{}
	statesGuard = &sync.RWMutex{}
)

func SturdyLogState(ctx context.Context, logger *logging.Logger, s *State) {
	logger = logger.Sublogger("sturdy")
	logger = logger.Sublogger(fmt.Sprintf("session.%s", s.Session.Version))
	statesGuard.RLock()
	ts, exists := statesMap[s.Session.Identifier]
	statesGuard.RUnlock()

	if exists {
		logger.Infof("%s: %s -> %s: %s", s.Session.Name, ts.Status, s.Status, time.Since(ts.Timestamp))
	}

	st := &stateTimestamp{
		Status:    s.Status,
		Timestamp: time.Now(),
	}
	labels, ok := sturdycontext.Labels(ctx)
	if ok {
		st.Labels = labels
	} else {
		st.Labels = make(map[string]string)
	}

	statesGuard.Lock()
	statesMap[s.Session.Identifier] = st
	statesGuard.Unlock()

	// The state is quite chatty, and we don't need to know all intermediate states.
	// Debouncing here to report the state at most once per second for each session.
	debouncesMx.Lock()

	// Nothing in queue, send right away
	if len(states) == 0 {
		states <- s
		debouncesMx.Unlock()
		return
	}

	// Debounce
	fn, ok := debounces[s.Session.Identifier]
	if !ok {
		debounces[s.Session.Identifier] = debounce.New(time.Second)
		fn = debounces[s.Session.Identifier]
	}
	debouncesMx.Unlock()

	fn(func() {
		states <- s
	})
}

type stateStatus struct {
	Identifier     string         `json:"identifier,omitempty"`
	Name           string         `json:"name,omitempty"`
	Status         string         `json:"status,omitempty"`
	AlphaURL       string         `json:"alphaUrl,omitempty"`
	BetaURL        string         `json:"betaUrl,omitempty"`
	AlphaConnected bool           `json:"alphaConnected"`
	BetaConnected  bool           `json:"betaConnected"`
	LastError      string         `json:"lastError,omitempty"`
	StagingStatus  receiverStatus `json:"stagingStatus,omitempty"`
	AlphaProblems  []problem      `json:"alphaProblems,omitempty"`
	BetaProblems   []problem      `json:"betaProblems,omitempty"`
	Paused         bool           `json:"paused"`
	SturdyVersion  string         `json:"sturdyVersion,omitempty"`
}

type problem struct {
	Path  string `json:"path,omitempty"`
	Error string `json:"error,omitempty"`
}

type receiverStatus struct {
	Path     string `json:"path,omitempty"`
	Received uint64 `json:"received,omitempty"`
	Total    uint64 `json:"total,omitempty"`
}

func problems(in []*core.Problem) []problem {
	res := make([]problem, len(in), len(in))
	for k, v := range in {
		res[k] = problem{
			Path:  v.Path,
			Error: v.Error,
		}
	}
	return res
}

func stateReporter() {
	for s := range states {
		if err := reportState(s); err != nil {
			log.Printf("failed to report state to Sturdy: %s", err)
		}
	}
}

func reportState(s *State) error {
	// Convert the State to the format that the Sturdy API is expecting.
	ss := &stateStatus{
		Identifier:     s.Session.Identifier,
		Name:           s.Session.Name,
		Status:         s.Status.String(),
		AlphaURL:       s.Session.Alpha.String(),
		BetaURL:        s.Session.Beta.String(),
		AlphaConnected: s.AlphaConnected,
		BetaConnected:  s.BetaConnected,
		LastError:      s.LastError,
		// TODO: Differentiate between Scan and Transition problems
		AlphaProblems: problems(s.AlphaScanProblems),
		BetaProblems:  problems(s.BetaScanProblems),
		Paused:        s.Session.Paused,
		SturdyVersion: SturdyVersion,
	}
	if s.StagingStatus != nil {
		ss.StagingStatus = receiverStatus{
			Path:     s.StagingStatus.Path,
			Received: s.StagingStatus.Received,
			Total:    s.StagingStatus.Total,
		}
	}

	ctx := sturdycontext.WithLabels(context.Background(), s.Session.Labels)
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second)
	defer cancelFunc()

	if err := api.Post(ctx, "/v3/mutagen/update-status", ss, &struct{}{}); err != nil {
		return fmt.Errorf("unexpected response: %w", err)
	}
	return nil
}
