package handler

import (
	"html/template"
	"net/http"

	"github.com/Tesohh/top100things/render"
)

func qerrFormat(qerr string) template.HTML {
	switch qerr {
	case "cred":
		return "Wrong email or password"
	case "pwlen":
		return "Invalid password"
	case "email":
		return "Invalid email"
	case "name":
		return "Invalid username"
	case "alrexists":
		return "User already exists"
	case "alrlogin":
		return template.HTML("You need to <a href=\"/logout\">logout</a> first")
	default:
		return ""
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	qerr := r.URL.Query().Get("err")
	serr := qerrFormat(qerr)

	render.Template(w, "login", map[string]any{"type": "login", "err": serr})
}

func Signup(w http.ResponseWriter, r *http.Request) {
	qerr := r.URL.Query().Get("err")
	serr := qerrFormat(qerr)

	render.Template(w, "login", map[string]any{"type": "signup", "err": serr})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "AccessToken",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
