package models

type Column struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Default     string   `json:"default"`
	Constraints []string `json:"constraints"`
}
