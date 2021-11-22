package model

import "time"

type SearchBase struct {
	Search string
	Limit  int
	Skip   int
}
type SearchBaseTime struct {
	Search    string     `json:"search"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	Limit     int        `json:"limit"`
	Skip      int        `json:"skip"`
}

type SearchMaintenance struct {
	State     string     `json:"state"`
	Limit     int        `json:"limit"`
	Skip      int        `json:"skip"`
}

type Records struct {
	Total   int         `json:"total"`
	Limit   int         `json:"limit"`
	Skip    int         `json:"skip"`
	Records interface{} `json:"records"`
}
