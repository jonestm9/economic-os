package fred

type ReleaseDate struct {
	ReleaseID   int    `json:"release_id"`
	ReleaseName string `json:"release_name"`
	Date        string `json:"date"`
}

type ReleaseDatesResponse struct {
	RealtimeStart string        `json:"realtime_start"`
	RealtimeEnd   string        `json:"realtime_end"`
	OrderBy       string        `json:"order_by"`
	SortOrder     string        `json:"sort_order"`
	Count         int           `json:"count"`
	Offset        int           `json:"offset"`
	Limit         int           `json:"limit"`
	ReleaseDates  []ReleaseDate `json:"release_dates"`
}

type Release struct {
	ID            int    `json:"id"`
	RealtimeStart string `json:"realtime_start"`
	RealtimeEnd   string `json:"realtime_end"`
	Name          string `json:"name"`
	PressRelease  bool   `json:"press_release"`
	Link          string `json:"link"`
}

type ReleaseResponse struct {
	RealtimeStart string    `json:"realtime_start"`
	RealtimeEnd   string    `json:"realtime_end"`
	Releases      []Release `json:"releases"`
}