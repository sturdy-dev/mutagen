package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	sturdy_context "github.com/mutagen-io/mutagen/pkg/sturdy/context"
)

var apiAddr string
var clientVersion = "development"

func init() {
	apiAddr = os.Getenv("STURDY_API_ADDR")
}

func getAPIAddr(ctx context.Context) string {
	if apiAddr != "" {
		return apiAddr
	}
	labels, ok := sturdy_context.Labels(ctx)
	if !ok {
		return apiAddr
	}
	proto := labels["sturdyApiProto"]
	host := labels["sturdyApiHost"]
	port := labels["sturdyApiHostPort"]
	prefix := labels["sturdyApiPrefix"]
	if port == "" {
		if prefix == "" {
			return fmt.Sprintf("%s://%s", proto, host)
		}
		return fmt.Sprintf("%s://%s/%s", proto, host, prefix)
	}
	if prefix == "" {
		return fmt.Sprintf("%s://%s:%s", proto, host, port)
	}
	return fmt.Sprintf("%s://%s:%s/%s", proto, host, port, prefix)
}

func Post(ctx context.Context, endpoint string, request, response interface{}) error {
	apiAddr := getAPIAddr(ctx)
	if apiAddr == "" {
		return fmt.Errorf("api address is not defined")
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, apiAddr+endpoint, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("could not make request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", clientVersion)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("unauthorized")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response: %w", err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("could not unmarshal response: %w", err)
	}

	return nil
}

func Get(ctx context.Context, endpoint string, response interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiAddr+endpoint, &bytes.Reader{})
	if err != nil {
		return fmt.Errorf("could not make request: %w", err)
	}
	req.Header.Add("User-Agent", clientVersion)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("unauthorized")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response: %w", err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("could not unmarshal response: %w", err)
	}

	return nil
}
