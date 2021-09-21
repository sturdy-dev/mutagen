package sturdy

import (
	"context"
	"fmt"

	"github.com/mutagen-io/mutagen/pkg/sturdy/api"
)

func ListIgnores(ctx context.Context, root string) ([]string, error) {
	_, viewID, err := ParseCodebaseViewPath(root)
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
