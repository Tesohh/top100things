package handler

import (
	"fmt"
	"net/http"

	"github.com/Tesohh/top100things/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, "home", "brodi")
	if err != nil {
		fmt.Fprint(w, "error loading the site")
	}
}
