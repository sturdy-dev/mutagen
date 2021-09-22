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
	if port == "" {
		return fmt.Sprintf("%s://%s", proto, host)
	}
	return fmt.Sprintf("%s://%s:%s", proto, host, port)
}

func Post(endpoint string, request, response interface{}) error {
	if apiAddr == "" {
		return fmt.Errorf("STURDY_API_ADDR is not defined")
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	resp, err := http.Post(apiAddr+endpoint, "application/json", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("could not make request: %w", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("unauthorized")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response: %w", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("could not unmarshal response: %w", err)
	}

	return nil
}

func Get(ctx context.Context, endpoint string, response interface{}) error {
	apiAddr = getAPIAddr(ctx)
	if apiAddr == "" {
		return fmt.Errorf("api address is not defined")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiAddr+endpoint, &bytes.Reader{})
	if err != nil {
		return fmt.Errorf("could not make request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not send request: %w", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("unauthorized")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response: %w", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("could not unmarshal response: %w", err)
	}

	return nil
}
