package model

import "time"

type Transaction struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	FromAccount int       `json:"from_account"`
	ToAccount   int       `json:"to_account"`
	Amount      float64   `json:"amount"`
}

type NewTransaction struct {
	CreatedAt   time.Time `json:"created_at"`
	FromAccount int       `json:"from_account"`
	ToAccount   int       `json:"to_account"`
	Amount      float64   `json:"amount"`
}
