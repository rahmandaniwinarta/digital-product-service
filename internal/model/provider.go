package model

import "time"

type Provider struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	CallbackURL string    `json:"callback_url"`
	CreatedAt   time.Time `json:"created_at"`
}
