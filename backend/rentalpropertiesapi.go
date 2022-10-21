package main

type RentalData struct {
	Id     float64   `json:"id"`
	Months []string  `json:"months"`
	Rents  []float64 `json:"rents"`
}
