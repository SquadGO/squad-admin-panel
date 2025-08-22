package handlers

import (
	"fmt"
	"net/http"

	"github.com/SquadGO/squad-admin-panel/internal/http/helpers"
	"github.com/SquadGO/squad-admin-panel/internal/models"
	"github.com/SquadGO/squad-admin-panel/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type AuthHandler interface {
	Auth(ctx *gin.Context)
	AuthSuccess(ctx *gin.Context)
}

type authHandler struct {
	userService service.UserService
}

func NewAuthHandler(userService service.UserService) AuthHandler {
	return &authHandler{
		userService: userService,
	}
}

func (a *authHandler) Auth(ctx *gin.Context) {
	redirectUrl := ctx.Query("redirect_url")

	if redirectUrl == "" {
		helpers.JsonError(ctx, http.StatusBadRequest, "Missing redirect_url param")
		return
	}

	ctx.SetCookie("redirect_url", redirectUrl, 300, "/", "", false, true)
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (a *authHandler) AuthSuccess(ctx *gin.Context) {
	redirectUrl, err := ctx.Cookie("redirect_url")
	if err != nil {
		helpers.JsonError(ctx, http.StatusBadRequest, "Missing redirect_url param")
		return
	}

	user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		helpers.JsonError(ctx, http.StatusForbidden, "Invalid steam")
		return
	}

	token, err := helpers.GenerateJWTToken(user.NickName, user.AvatarURL, user.UserID)
	if err != nil {
		helpers.JsonError(ctx, http.StatusForbidden, "Failed generate token")
		return
	}

	a.userService.CreateUser(ctx, models.CreateUser{SteamID: user.UserID, Name: user.NickName, Avatar: &user.AvatarURL})

	ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s?token=%s", redirectUrl, token))
}
