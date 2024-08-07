package domain

import "time"

type Todo struct {
	ID       string    `db:"id" json:"id"`
	Title    string    `db:"title" json:"title"`
	ActiveAt time.Time `db:"active_at" json:"active_at"`
	Status   bool      `db:"status" json:"status"`
}
