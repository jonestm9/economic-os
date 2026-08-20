package config

type Series struct {
	ReleaseID int
	Source    string
	SeriesID  string
	Name      string
}

var TrackedSeries = []Series{
	{
		ReleaseID: 50,
		Source:    "BLS",
		SeriesID:  "LNS14000000",
		Name:      "Unemployment Rate",
	},
}