package auth

import (
	"context"
	"net/http"

	"github.com/Tesohh/top100things/sb"
	"github.com/nedpals/supabase-go"
)

func User(r *http.Request) *supabase.User {
	ac, aerr := r.Cookie("AccessToken")
	rc, rerr := r.Cookie("RefreshToken")
	_ = rc
	if aerr != nil || rerr != nil { // cookies don't exist == timed out or logged out
		return nil
	}

	user, err := sb.SB.Auth.User(context.Background(), ac.Value)
	if user == nil || err != nil { // the user cannot be retrieved, probably due to the token being timed out
		ad, err := sb.SB.Auth.RefreshUser(context.Background(), ac.Value, rc.Value)
		if err != nil {
			return nil
		}
		user = &ad.User
	}

	return user
}
