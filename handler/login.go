package handler

import (
	"net/http"

	"github.com/Tesohh/top100things/render"
)

func Login(w http.ResponseWriter, r *http.Request) {
	qerr := r.URL.Query().Get("err")
	var serr string
	switch qerr {
	case "cred":
		serr = "Wrong email or password"
	case "pwlen":
		serr = "Invalid password"
	case "email":
		serr = "Invalid email"
	default:
		serr = ""
	}
	render.Template(w, "login", map[string]any{"type": "login", "err": serr})
}
