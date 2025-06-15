package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yohcop/openid-go"
)

type AuthHandler interface {
	Auth(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authHandler struct{}

var (
	nonceStore     = openid.NewSimpleNonceStore()
	discoveryCache = openid.NewSimpleDiscoveryCache()
)

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

func (as *authHandler) Auth(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	returnURL := "http://localhost:3001/login"

	authURL, err := openid.RedirectURL(
		"https://steamcommunity.com/openid",
		returnURL,
		returnURL,
	)
	if err != nil {
		http.Error(w, "Failed to generate Steam auth URL", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, authURL, http.StatusFound)
}

func (as *authHandler) Login(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request

	fullURL := "http://localhost:3001" + r.URL.String()
	id, err := openid.Verify(fullURL, discoveryCache, nonceStore)
	if err != nil {
		http.Error(w, "OpenID verification failed", http.StatusUnauthorized)
		return
	}

	steamID := id[len("https://steamcommunity.com/openid/id/"):]
	fmt.Fprintf(w, "User authenticated with SteamID: %s", steamID)
}
