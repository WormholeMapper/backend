package esiclient

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/WormholeMapper/whmcfg"
	"github.com/antihax/goesi"
)

// GenerateState generates a 16-byte base64 encoded string, for use with the ESI API.
func GenerateState() (string, error) {
	stateBuf := make([]byte, 16)
	_, err := rand.Read(stateBuf)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(stateBuf), nil
}

// SetState sets a short-lived cookie containing the ESI OAuth state value on the client.
func SetState(w http.ResponseWriter, state string) {
	http.SetCookie(w, &http.Cookie{
		Name:   "esi-oauth-state",
		Value:  state,
		MaxAge: 300,
	})
}

type key string

// WithSSOAuthenticator is middleware providing access to the EVE SSO authentication
// endpoint generator, from antihax/goesi.
func WithSSOAuthenticator(next http.Handler, config whmcfg.AppConfig) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authenticator := goesi.NewSSOAuthenticatorV2(
				http.DefaultClient,
				config.Esi.ClientID,
				config.Esi.SecretKey,
				config.Esi.CallbackURL,
				nil,
			)
			// Set the authenticator object as a key:value pair on the request context.
			ctx := context.WithValue(r.Context(), key("authenticator"), authenticator)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

// SSOAuth initialises an ESI OAuth session.
// This endpoint must be wrapped with the WithSSOAuthenticator middleware.
func SSOAuth(w http.ResponseWriter, r *http.Request) (int, error) {
	// Generate the state and store it as a cookie on the response.
	state, err := GenerateState()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	SetState(w, state)
	// Try to get the authenticator object from the context.
	// TODO: try to fix the linter error here: the authenticator object contains a sync.Mutex,
	// which we aren't really supposed to copy here. Possible there's no harm, but needs testing.
	authenticator, ok := r.Context().Value("authenticator").(goesi.SSOAuthenticator)
	if !ok {
		return http.StatusInternalServerError, errors.New("unable to get authenticator from request context")
	}
	http.Redirect(w, r, authenticator.AuthorizeURL(state, false, nil), http.StatusMovedPermanently)
	return http.StatusMovedPermanently, nil
}
