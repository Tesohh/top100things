package model

import "time"

type Model struct {
	ID        int8      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
