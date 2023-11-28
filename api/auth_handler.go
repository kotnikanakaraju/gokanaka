// api/auth_handler.go
package api

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

var (
	oauthConfig *oauth2.Config
	oauthStateString="random"
)

func init() {
    oauthConfig = &oauth2.Config{
        ClientID:     "000000",
        ClientSecret: "999999",
        RedirectURL:  "http://localhost:8080/auth/callback",
        Endpoint: oauth2.Endpoint{
            AuthURL:  "https://accounts.google.com/o/oauth2/auth",
            TokenURL: "https://accounts.google.com/o/oauth2/token",
        },
        Scopes: []string{"profile", "email"},
    }
}

func AuthenticateUser(r *fiber.Ctx) error {
    url := oauthConfig.AuthCodeURL(oauthStateString)
    return r.Redirect(url, http.StatusTemporaryRedirect)

}

func HandleOAuthCallback(r *fiber.Ctx) error {
    state := r.Query("state")
    if state != oauthStateString {
        return r.SendStatus(http.StatusUnauthorized)
    }

    code := r.Query("code")
    token, err := oauthConfig.Exchange(r.Context(), code)
    if err != nil {
        return r.SendStatus(http.StatusInternalServerError)
    }

    fmt.Println(token)
    

    
    return r.SendString("Authentication successful!",)
}
