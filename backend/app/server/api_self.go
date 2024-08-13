package server

import (
	"errors"
	"strconv"

	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/utils/libx/httpx"
	"github.com/bornjre/turnix/backend/xtypes/models"
	"github.com/gin-gonic/gin"
)

func (a *Server) selfInfo(claim *signer.AccessClaim, _ *gin.Context) (any, error) {

	usr, err := a.db.GetUser(claim.UserId)
	if err != nil {
		return nil, err
	}

	usr.Password = ""

	return usr, nil
}

type changePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

var errIncorrectOldPassword = errors.New("err invalid old password")

func (a *Server) selfChangePassword(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

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

	return httpx.MessageOk, nil
}

func (a *Server) selfUsers(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	return a.db.ListUserByOwner(claim.UserId)
}

func (a *Server) selfAddUser(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	data := &models.User{}

	err := ctx.Bind(data)
	if err != nil {
		return nil, err
	}

	data.OwnerUserId = claim.UserId

	return a.db.AddUser(data)
}

var errNotAllowed = errors.New("err Not Allowed")

func (a *Server) selfGetUser(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	usr, err := a.db.GetUser(claim.UserId)
	if err != nil {
		return nil, err
	}

	if usr.OwnerUserId != claim.UserId {
		return nil, errNotAllowed
	}

	usr.Password = ""

	return usr, nil
}

func (a *Server) selfUpdateUser(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

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

	if usr.OwnerUserId == claim.UserId {
		return nil, errNotAllowed
	}

	// fixme => only allow some fields to be updated

	err = a.db.UpdateUser(uid, data)
	if err != nil {
		return nil, err
	}

	return httpx.MessageOk, nil
}

func (a *Server) selfDeleteUser(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	uid, err := strconv.ParseInt(ctx.Param("uid"), 10, 64)
	if err != nil {
		return nil, err
	}

	usr, err := a.db.GetUser(claim.UserId)
	if err != nil {
		return nil, err
	}

	if usr.OwnerUserId == claim.UserId {
		return nil, errNotAllowed
	}

	err = a.db.DeleteUser(uid)
	if err != nil {
		return nil, err
	}

	return httpx.MessageOk, nil
}
