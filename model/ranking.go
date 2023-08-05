package model

type Ranking struct {
	Model
	UserID      string `json:"user_id"`
	Thing       Thing  `json:"thing_id"`
	Position    int    `json:"position"`
	Description string `json:"description"`
}
