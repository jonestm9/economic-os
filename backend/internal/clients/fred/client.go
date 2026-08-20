package fred

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const fredAPIBaseURL = "https://api.stlouisfed.org/fred"

type Client struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient(apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		apiKey: apiKey,
	}
}

func (c *Client) get(endpoint string, params url.Values, result any) error {
	params.Set("api_key", c.apiKey)
	params.Set("file_type", "json")

	requestURL := fmt.Sprintf(
		"%s/%s?%s",
		fredAPIBaseURL,
		endpoint,
		params.Encode(),
	)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return fmt.Errorf("creating FRED request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("requesting FRED API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"FRED API returned status code %d",
			resp.StatusCode,
		)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("decoding FRED response: %w", err)
	}

	return nil
}

func (c *Client) GetReleaseDates() ([]ReleaseDate, error) {
	params := url.Values{}
	params.Set("include_release_dates_with_no_data", "true")
	params.Set("order_by", "release_date")
	params.Set("sort_order", "asc")

	var response ReleaseDatesResponse

	if err := c.get("releases/dates", params, &response); err != nil {
		return nil, fmt.Errorf("getting release dates: %w", err)
	}

	return response.ReleaseDates, nil
}

func (c *Client) GetRelease(releaseID int) (*Release, error) {
	params := url.Values{}
	params.Set("release_id", fmt.Sprintf("%d", releaseID))

	var response ReleaseResponse

	if err := c.get("release", params, &response); err != nil {
		return nil, fmt.Errorf(
			"getting FRED release %d: %w",
			releaseID,
			err,
		)
	}

	if len(response.Releases) == 0 {
		return nil, fmt.Errorf(
			"FRED returned no release for ID %d",
			releaseID,
		)
	}

	return &response.Releases[0], nil
}