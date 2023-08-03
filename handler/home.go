package handler

import (
	"fmt"
	"net/http"

	"github.com/Tesohh/top100things/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	err := template.Parse(w, "home", "brodi")
	if err != nil {
		fmt.Fprint(w, "error loading the site")
	}
}
