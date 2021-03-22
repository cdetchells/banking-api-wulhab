package model

type Account struct {
	ID       int     `json:"id"`
	Customer int     `json:"customer"`
	Type     int8    `json:"type"`
	Balance  float64 `json:"balance"`
}

type NewAccount struct {
	Customer int     `json:"customer"`
	Type     int8    `json:"type"`
	Balance  float64 `json:"balance"`
}
