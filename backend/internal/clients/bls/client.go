package bls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const blsURL = "https://api.bls.gov/publicAPI/v2/timeseries/data/"

func NewClient(httpClient *http.Client, apiKey string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &Client{
		HTTPClient: httpClient,
		APIKey:     apiKey,
	}
}

// FetchData sends a POST request to retrieve time-series data.
func (c *Client) FetchData(seriesIDs []string, startYear, endYear string) (*Response, error) {
	reqPayload := Request{
		SeriesID:        seriesIDs,
		StartYear:       startYear,
		EndYear:         endYear,
		RegistrationKey: c.APIKey,
	}

	jsonBytes, err := json.Marshal(reqPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		blsURL,
		bytes.NewBuffer(jsonBytes),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var blsResp Response

	err = json.NewDecoder(resp.Body).Decode(&blsResp)
	if err != nil {
		return nil, fmt.Errorf("failed decoding response: %w", err)
	}

	if blsResp.Status != "REQUEST_SUCCEEDED" {
		return &blsResp, fmt.Errorf(
			"BLS request failed: %s",
			blsResp.Message,
		)
	}

	return &blsResp, nil
}