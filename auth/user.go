package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/Tesohh/top100things/sb"
	"github.com/nedpals/supabase-go"
)

func User(r *http.Request) *supabase.User {
	ac, aerr := r.Cookie("AccessToken")
	rc, rerr := r.Cookie("RefreshToken")
	if aerr == nil && rerr == nil { // if the cookie exists
		// user, err := sb.SB.Auth.User(context.Background(), ac.Value)
		ad, err := sb.SB.Auth.RefreshUser(context.Background(), ac.Value, rc.Value)
		if err == nil {
			return &ad.User
		} else {
			log.Println("user: " + err.Error())
		}
	}
	return nil

}
