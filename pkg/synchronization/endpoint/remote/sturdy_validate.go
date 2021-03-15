package remote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type sturdyValidateViewer func(codebaseID, viewID string) error

func sturdyValidateRoot(root string, sturdyValFunc sturdyValidateViewer) error {
	root = path.Clean(root)
	parts := strings.Split(root, "/")

	if len(parts) != 4 {
		return fmt.Errorf("unknown number of paths")
	}

	if parts[0] != "" || parts[1] != "repos" {
		return fmt.Errorf("invalid path")
	}

	err := sturdyValFunc(parts[2], parts[3])
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

	apiAddr := os.Getenv("STURDY_API_ADDR")
	if len(apiAddr) == 0 {
		return fmt.Errorf("could not get STURDY_API_ADDR")
	}

	userID := os.Getenv("STURDY_AUTHENTICATED_USER_ID")
	if len(apiAddr) == 0 {
		log.Println("could not get STURDY_AUTHENTICATED_USER_ID")
		return fmt.Errorf("could not get STURDY_AUTHENTICATED_USER_ID")
	}

	data, err := json.Marshal(validateRequest{
		UserID:     userID,
		ViewID:     viewID,
		CodebaseID: codebaseID,
	})
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	resp, err := http.Post(apiAddr+"/v3/mutagen/validate-view", "application/json", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("could not make request: %w", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("unauthorized")
	}

	return nil
}
