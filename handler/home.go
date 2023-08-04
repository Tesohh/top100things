package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Tesohh/top100things/render"
	"github.com/Tesohh/top100things/sb"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]any)
	c, err := r.Cookie("AccessToken")
	if err == nil { // if the cookie exists
		user, err := sb.SB.Auth.User(context.Background(), c.Value)
		if err != nil {
			fmt.Println("NOT logged in")
		} else {
			data["user"] = user
		}
	}

	err = render.Template(w, "home", data)
	if err != nil {
		fmt.Fprint(w, err)
	}
}
