package model

import (
	"fmt"
)

type Thing struct {
	Model
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	AuthorID string `json:"author"`
}

func (t Thing) String() string {
	return fmt.Sprintf("%v [%v] by %v", t.Name, t.ImageURL, t.AuthorID)
}
