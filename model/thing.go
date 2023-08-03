package model

import (
	"fmt"
)

type Thing struct {
	Model
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func (t Thing) String() string {
	return fmt.Sprintf("%v [%v] at %v", t.Name, t.ImageURL, t.CreatedAt)
}
