package models

type ApiDetails struct {
	Path       string `json:"path"`
	MinVersion int    `json:"minVersion"`
	Maxversion int    `json:"maxVersion"`
}

type ApiInfo map[string]ApiDetails
