package model

type Address struct {
	Line1   string `json:"line_1"`
	Line2   string `json:"line_2,omitempty"`
	Line3   string `json:"line_3,omitempty"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}
