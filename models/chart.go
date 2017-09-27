package models

type (
	//Chart : represent data for chart
	Chart struct {
		Labels       []string  `json:"labels"`
		Custody      []float32 `json:"custody"`
		Accomplished []float32 `json:"accomplished"`
		Meta         []float32 `json:"meta"`
	}
)
