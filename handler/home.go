package handler

import (
	"fmt"
	"net/http"

	"github.com/Tesohh/top100things/auth"
	"github.com/Tesohh/top100things/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, "home", map[string]any{"user": auth.User(r)})
	if err != nil {
		fmt.Fprint(w, err)
	}
}
