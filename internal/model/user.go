package model

import "time"

type User struct {
	ID        int64     `json : "id"`
	Name      string    `json : "name"`
	Email     string    `json : "email"`
	Balance   int64     `json : "balance"`
	CreatedAt time.Time `json : "created_at"`
}
