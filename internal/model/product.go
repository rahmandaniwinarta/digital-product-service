package model

import "time"

type Product struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Price      int64     `json:"price"`
	ProviderID int64     `json:"provider_id"`
	CreatedAt  time.Time `json:"created_at"`
}
