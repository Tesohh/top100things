package form

import (
	"context"
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"github.com/Tesohh/top100things/sb"
	"github.com/nedpals/supabase-go"
)

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if _, err := mail.ParseAddress(email); err != nil {
		http.Redirect(w, r, "/login?err=email", http.StatusSeeOther)
		return
	} else if len(password) < 8 {
		http.Redirect(w, r, "/login?err=pwlen", http.StatusSeeOther)
		return
	}

	ctx := context.Background()
	user, err := sb.SB.Auth.SignIn(ctx, supabase.UserCredentials{Email: email, Password: password})
	if err != nil {
		http.Redirect(w, r, "/login?err=cred", http.StatusSeeOther)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "AccessToken",
		Value:    user.AccessToken,
		Expires:  time.Now().Add(365 * 24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	// username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	_, cookieerr := r.Cookie("AccessToken")
	if cookieerr == nil { // if theres already a cookie
		http.Redirect(w, r, "/signup?err=alrlogin", http.StatusSeeOther)
		return
	} else if _, err := mail.ParseAddress(email); err != nil {
		http.Redirect(w, r, "/signup?err=email", http.StatusSeeOther)
		return
	} else if len(password) < 8 {
		http.Redirect(w, r, "/signup?err=pwlen", http.StatusSeeOther)
		return
	}

	ctx := context.Background()
	_, err := sb.SB.Auth.SignUp(ctx, supabase.UserCredentials{Email: email, Password: password})
	if err != nil {
		if err.Error() == "User already registered" { // i know this is garbage code but it works
			http.Redirect(w, r, "/signup?err=alrexists", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/signup?err=cred", http.StatusSeeOther)
		}
		return
	}

	tokenholder, err := sb.SB.Auth.SignIn(ctx, supabase.UserCredentials{Email: email, Password: password})
	if err != nil {
		http.Redirect(w, r, "/signup?err=cred", http.StatusSeeOther)
		fmt.Println(err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "AccessToken",
		Value:    tokenholder.AccessToken,
		Expires:  time.Now().Add(365 * 24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
