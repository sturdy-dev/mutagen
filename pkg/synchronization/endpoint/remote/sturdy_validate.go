package remote

import (
	"fmt"
	"github.com/mutagen-io/mutagen/pkg/sturdy"
	"github.com/mutagen-io/mutagen/pkg/sturdy/api"
	"log"
	"os"
)

type sturdyValidateViewer func(codebaseID, viewID string) error

func sturdyValidateRoot(root string, sturdyValFunc sturdyValidateViewer) error {
	codebaseID, viewID, err := sturdy.ParseCodebaseViewPath(root)
	if err != nil {
		return fmt.Errorf("invalid path")
	}

	err = sturdyValFunc(codebaseID, viewID)
	if err != nil {
		return err
	}

	return nil
}

func sturdyApiValidateRoot(codebaseID, viewID string) error {
	type validateRequest struct {
		CodebaseID string `json:"codebase_id"`
		ViewID     string `json:"view_id"`
		UserID     string `json:"user_id"`
	}

	userID := os.Getenv("STURDY_AUTHENTICATED_USER_ID")
	if len(userID) == 0 {
		log.Println("could not get STURDY_AUTHENTICATED_USER_ID")
		return fmt.Errorf("could not get STURDY_AUTHENTICATED_USER_ID")
	}

	var res struct{}
	err := api.Request("/v3/mutagen/validate-view", validateRequest{
		UserID:     userID,
		ViewID:     viewID,
		CodebaseID: codebaseID,
	}, &res)
	if err != nil {
		return fmt.Errorf("validate view failed: %w", err)
	}

	return nil
}
