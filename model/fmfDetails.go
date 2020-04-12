package model

type FMFDetails struct {
	FilterType  string `json:"filter_type,omitempty"`
	FilterID    string `json:"filter_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
