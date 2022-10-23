package telegraph

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Body struct {
	// Ok: if true, request was successful, and result can be found in the Result field.
	// If false, error can be explained in Error field.
	Ok bool `json:"ok"`
	// Error: contains a human readable description of the error result.
	Error string `json:"error"`
	// Result: result of requests (if Ok)
	Result json.RawMessage `json:"result"`
}

func InvokeRequest(method string, params url.Values) (json.RawMessage, error) {
	r, err := http.NewRequest(http.MethodPost, "https://api.telegra.ph/"+method, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build GET request to %s: %w", method, err)
	}
	r.URL.RawQuery = params.Encode()
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GET request to %s: %w", method, err)
	}
	defer resp.Body.Close()

	var b Body

	if err = json.NewDecoder(resp.Body).Decode(&b); err != nil {
		return nil, fmt.Errorf("failed to decode GET request to %s: %w", method, err)
	}
	if !b.Ok {
		return nil, fmt.Errorf("failed to %s: %s", method, b.Error)
	}
	return b.Result, nil
}
