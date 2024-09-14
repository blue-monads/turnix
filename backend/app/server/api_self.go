package server

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bornjre/turnix/backend/controller/self"
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

// messages

func (a *Server) listUserMessages(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	count, _ := strconv.ParseInt(ctx.Query("count"), 10, 64)
	cursor, _ := strconv.ParseInt(ctx.Query("cursor"), 10, 64)
	return a.cSelf.ListUserMessages(claim.UserId, count, cursor)
}

func (a *Server) messageUser(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	data := &self.MessegeRequest{}

	err := ctx.Bind(data)
	if err != nil {
		return nil, err
	}

	uid, _ := strconv.ParseInt(ctx.Param("uid"), 10, 64)
	if uid == 0 {
		return nil, fmt.Errorf("user id is required")
	}

	err = a.cSelf.SendUserMessage(claim.UserId, uid, data)
	return nil, err
}

// self files

func (a *Server) listSelfFiles(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	path := ctx.Query("path")

	return a.cSelf.ListSelfFiles(claim.UserId, path)
}

func (a *Server) addSelfFolder(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	data := CreateFolder{}
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	return a.cSelf.AddSelfFolder(claim.UserId, data.Path, data.Name)
}

func (a *Server) addSelfFile(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	name := ctx.Query("name")
	path := ctx.Query("path")

	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	return a.cSelf.AddSelfFile(claim.UserId, name, path, ctx.Request.ContentLength, ctx.Request.Body)
}

func (a *Server) getSelfFile(ctx *gin.Context) {

	claim, err := a.withAccess(ctx)
	if err != nil {
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err = a.cSelf.GetSelfFile(claim.UserId, id, ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

}

func (a *Server) removeSelfFile(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err := a.cSelf.DeleteSelfFile(claim.UserId, id)

	return nil, err
}

// file shares

func (a *Server) listSelfFileShares(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	fileId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	return a.cSelf.ListFileShares(claim.UserId, fileId)
}

func (a *Server) addSelfFileShare(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	fileId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	return a.cSelf.AddFileShare(claim.UserId, fileId)
}

func (a *Server) deleteSelfFileShare(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	err := a.cSelf.DeleteFileShare(claim.UserId, ctx.Param("id"))
	return nil, err
}
