package remote

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/mutagen-io/mutagen/pkg/sturdy"
	"github.com/mutagen-io/mutagen/pkg/sturdy/api"
)

type sturdyValidateViewer func(ctx context.Context, codebaseID, viewID string, isNewConnection bool) error

func sturdyValidateRoot(ctx context.Context, root string, sturdyValFunc sturdyValidateViewer) error {
	codebaseID, viewID, err := sturdy.ParseCodebaseViewPath(root)
	if err != nil {
		return fmt.Errorf("invalid path")
	}

	if err := sturdyValFunc(ctx, codebaseID, viewID, true); err != nil {
		return err
	}

	return nil
}

func sturdyApiValidateRoot(ctx context.Context, codebaseID, viewID string, isNewConnection bool) error {
	type validateRequest struct {
		CodebaseID      string `json:"codebase_id"`
		ViewID          string `json:"view_id"`
		UserID          string `json:"user_id"`
		IsNewConnection bool   `json:"is_new_connection"`
	}

	userID := os.Getenv("STURDY_AUTHENTICATED_USER_ID")
	if len(userID) == 0 {
		log.Println("could not get STURDY_AUTHENTICATED_USER_ID")
		return fmt.Errorf("could not get STURDY_AUTHENTICATED_USER_ID")
	}

	var res struct{}
	if err := api.Post(ctx, "/v3/mutagen/validate-view", validateRequest{
		UserID:          userID,
		ViewID:          viewID,
		CodebaseID:      codebaseID,
		IsNewConnection: isNewConnection,
	}, &res); err != nil {
		return fmt.Errorf("validate view failed: %w", err)
	}

	return nil
}

func sturdyPingView(ctx context.Context, root string, sturdyValFunc sturdyValidateViewer, done chan bool) error {
	codebaseID, viewID, err := sturdy.ParseCodebaseViewPath(root)
	if err != nil {
		return fmt.Errorf("invalid path")
	}

	// Interval with jitter [25s, 33s]
	r := rand.Intn(8000)
	tick := time.Tick(time.Second*25 + time.Millisecond*time.Duration(r))

	for {
		select {
		case <-done:
			return nil
		case <-tick:
			if err := sturdyValFunc(ctx, codebaseID, viewID, false); err != nil {
				return err
				log.Printf("failed to ping view: %s", err)
			}
		}
	}
}
