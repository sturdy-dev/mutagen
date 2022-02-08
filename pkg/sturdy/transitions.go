package sturdy

import (
	"context"
	"fmt"

	"github.com/mutagen-io/mutagen/pkg/sturdy/api"
)

type SyncTransitionsRequest struct {
	Paths      []string `json:"paths"`
	CodebaseID string   `json:"codebase_id"`
	ViewID     string   `json:"view_id"`
}

func SyncTransitions(ctx context.Context, root string, paths []string) error {
	codebaseID, viewID, err := ParseCodebaseViewPath(root)
	if err != nil {
		return err
	}

	var res struct{}
	if err := api.Post(ctx, "/v3/mutagen/sync-transitions", SyncTransitionsRequest{
		Paths:      paths,
		CodebaseID: codebaseID,
		ViewID:     viewID,
	}, &res); err != nil {
		return fmt.Errorf("failed to sync transitions: %w", err)
	}
	return nil
}
