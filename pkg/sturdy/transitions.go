package sturdy

import (
	"fmt"
	"github.com/mutagen-io/mutagen/pkg/sturdy/api"
)

type SyncTransitionsRequest struct {
	Paths      []string `json:"paths"`
	CodebaseID string   `json:"codebase_id"`
	ViewID     string   `json:"view_id"`
}

func SyncTransitions(root string, paths []string) error {
	codebaseID, viewID, err := ParseCodebaseViewPath(root)
	if err != nil {
		return err
	}

	var res struct{}
	err = api.Request("/v3/mutagen/sync-transitions", SyncTransitionsRequest{
		Paths:      paths,
		CodebaseID: codebaseID,
		ViewID:     viewID,
	}, &res)
	if err != nil {
		return fmt.Errorf("failed to sync transitions: %w", err)
	}
	return nil
}
