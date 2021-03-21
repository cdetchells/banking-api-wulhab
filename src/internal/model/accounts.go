package model

type Account struct {
	ID       int32   `json:"id"`
	Customer int32   `json:"customer"`
	Type     int8    `json:"type"`
	Balance  float64 `json:"balance"`
}
