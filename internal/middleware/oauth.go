package middleware

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "12345678"
	maxAge = 86400 * 30
	isProd = true
)

func NewOAuth() {
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd
	gothic.Store = store

	goth.UseProviders(
		google.New(
            os.Getenv("GOOGLE_CLIENT_ID"),
            os.Getenv("GOOGLE_CLIENT_SECRET"),
			"http://127.0.0.1:8080/auth/google/callback",
		),
        facebook.New(
            os.Getenv("FACEBOOK_CLIENT_ID"),
            os.Getenv("FACEBOOK_CLIENT_SECRET"),
            "http://127.0.0.1:8080/auth/facebook/callback",
        ),
	)
}
