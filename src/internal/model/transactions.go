package model

import "time"

type Transactions struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	FromAccount string    `json:"from_account"`
	ToAccount   string    `json:"to_account"`
	Amount      float64   `json:"amount"`
}
