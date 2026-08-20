package bls

import (
	"net/http"
)

//todo eventually split between observation and 
// Request represents the structured payload sent to the BLS API v2.
type Request struct {
	SeriesID        []string `json:"seriesid"`
	StartYear       string   `json:"startyear,omitempty"`
	EndYear         string   `json:"endyear,omitempty"`
	RegistrationKey string   `json:"registrationkey,omitempty"`
	Catalog         bool     `json:"catalog,omitempty"`
	Calculations    bool     `json:"calculations,omitempty"`
	AnnualAverage   bool     `json:"annualaverage,omitempty"`
}

// Footnote contains any explanatory notes appended to a data point.
type Footnote struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

// DataPoint represents an individual time-series entry returned by the BLS.
type DataPoint struct {
	Year       string     `json:"year"`
	Period     string     `json:"period"`
	PeriodName string     `json:"periodName"`
	Value      string     `json:"value"`
	Footnotes  []Footnote `json:"footnotes"`
}

// SeriesData represents a specific metric time-series sequence.
type SeriesData struct {
	SeriesID string      `json:"seriesID"`
	Data     []DataPoint `json:"data"`
}

// ResultsData acts as the wrapper holding all retrieved time-series data.
type ResultsData struct {
	Series []SeriesData `json:"series"`
}

// Response represents the full JSON document returned by the BLS API v2.
type Response struct {
	Status       string      `json:"status"`
	ResponseTime int         `json:"responseTime"`
	Message      []string    `json:"message"`
	Results      ResultsData `json:"results"`
}

// Client manages configurations and HTTP lifecycles for talking to the BLS API.
type Client struct {
	HTTPClient *http.Client
	APIKey     string
}