package app

import (
	"errors"
	"strconv"

	"github.com/bornjre/trunis/backend/database"
	"github.com/bornjre/trunis/backend/token"
	"github.com/gin-gonic/gin"
)

func (a *App) selfInfo(claim *token.AccessClaim, _ *gin.Context) (any, error) {

	usr, err := a.db.GetUser(claim.UserId)
	if err != nil {
		return nil, err
	}

	usr.AccessToken = ""
	usr.Password = ""

	return usr, nil
}

type changePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

var errIncorrectOldPassword = errors.New("err invalid old password")

func (a *App) selfChangePassword(claim *token.AccessClaim, ctx *gin.Context) (any, error) {

	data := &changePasswordRequest{}
	err := ctx.Bind(data)
	if err != nil {
		return nil, err
	}

	usr, err := a.db.GetUser(claim.UserId)
	if err != nil {
		return nil, err
	}

	if usr.Password != data.NewPassword {
		return nil, errIncorrectOldPassword
	}

	return MessageOk, nil
}

func (a *App) selfUsers(claim *token.AccessClaim, ctx *gin.Context) (any, error) {
	return a.db.ListUserByOwner(claim.UserId)
}

func (a *App) selfAddUser(claim *token.AccessClaim, ctx *gin.Context) (any, error) {

	data := &database.User{}

	err := ctx.Bind(data)
	if err != nil {
		return nil, err
	}

	data.AccessToken = ""
	data.OwnedBy = claim.UserId

	return a.db.AddUser(data)
}

var errNotAllowed = errors.New("err Not Allowed")

func (a *App) selfGetUser(claim *token.AccessClaim, ctx *gin.Context) (any, error) {

	usr, err := a.db.GetUser(claim.UserId)
	if err != nil {
		return nil, err
	}

	if usr.OwnedBy != claim.UserId {
		return nil, errNotAllowed
	}

	usr.AccessToken = ""
	usr.Password = ""

	return usr, nil
}

func (a *App) selfUpdateUser(claim *token.AccessClaim, ctx *gin.Context) (any, error) {

	data := make(map[string]any)
	ctx.Bind(&data)

	uid, err := strconv.ParseInt(ctx.Param("uid"), 10, 64)
	if err != nil {
		return nil, err
	}

	usr, err := a.db.GetUser(claim.UserId)
	if err != nil {
		return nil, err
	}

	if usr.OwnedBy == claim.UserId {
		return nil, errNotAllowed
	}

	// fixme => only allow some fields to be updated

	err = a.db.UpdateUser(uid, data)
	if err != nil {
		return nil, err
	}

	return MessageOk, nil
}

func (a *App) selfDeleteUser(claim *token.AccessClaim, ctx *gin.Context) (any, error) {

	uid, err := strconv.ParseInt(ctx.Param("uid"), 10, 64)
	if err != nil {
		return nil, err
	}

	usr, err := a.db.GetUser(claim.UserId)
	if err != nil {
		return nil, err
	}

	if usr.OwnedBy == claim.UserId {
		return nil, errNotAllowed
	}

	err = a.db.DeleteUser(uid)
	if err != nil {
		return nil, err
	}

	return MessageOk, nil
}
