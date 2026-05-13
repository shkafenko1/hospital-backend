package models

type Ward struct {
	ID       int64  `json:"id"`
	Number   int64  `json:"number"`
	Floor    int64  `json:"floor"`
	Capacity int64  `json:"capacity"`
	Price    int64  `json:"price"`
	Comfort  string `json:"comfort"`
	Busy     bool   `json:"busy"`
}
