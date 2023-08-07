package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/Tesohh/top100things/sb"
	"github.com/nedpals/supabase-go"
)

func User(r *http.Request) *supabase.User {
	c, err := r.Cookie("AccessToken")
	if err == nil { // if the cookie exists
		user, err := sb.SB.Auth.User(context.Background(), c.Value)
		if err == nil {
			return user
		} else {
			log.Println("user: " + err.Error())
		}
	}
	return nil

}
