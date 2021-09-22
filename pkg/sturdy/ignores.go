package sturdy

import (
	"context"
	"fmt"

	"github.com/mutagen-io/mutagen/pkg/sturdy/api"
	sturdy_context "github.com/mutagen-io/mutagen/pkg/sturdy/context"
)

func ListIgnores(ctx context.Context, root string) ([]string, error) {
	viewID, err := viewID(ctx, root)
	if err != nil {
		return nil, err
	}

	var res struct {
		Ignores []string `json:"ignores"`
	}

	url := fmt.Sprintf("/v3/mutagen/views/%s/ignores", viewID)
	if err := api.Get(ctx, url, &res); err != nil {
		return nil, err
	}

	return res.Ignores, nil
}

func viewID(ctx context.Context, root string) (string, error) {
	if _, viewID, err := ParseCodebaseViewPath(root); err == nil {
		// On the server-side, we can parse view id from the dir root.
		return viewID, nil
	}

	if labels, ok := sturdy_context.Labels(ctx); ok {
		// For the client-side sturdy-sync (mutagen), labels must be in the context.
		return labels["sturdyViewId"], nil
	}

	return "", fmt.Errorf("failed to fetch view id")
}
