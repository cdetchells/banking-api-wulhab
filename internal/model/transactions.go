package model

import "time"

type Transaction struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	FromAccount int       `json:"fromAccount"`
	ToAccount   int       `json:"toAccount"`
	Amount      float64   `json:"amount"`
}

type NewTransaction struct {
	CreatedAt time.Time `json:"createdAt"`
	ToAccount int       `json:"toAccount"`
	Amount    float64   `json:"amount"`
}
