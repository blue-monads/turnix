package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bornjre/turnix/backend/utils/libx/httpx"
	"github.com/bornjre/turnix/backend/xtypes/models"
	_ "github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"

	"github.com/bornjre/turnix/backend/services/signer"
)

func (a *Server) signUpDirect(ctx *gin.Context) {

	data := &models.User{}
	err := ctx.BindJSON(data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	if data.Email == "" || data.Name == "" || data.Password == "" {
		return
	}

	data.IsEmailVerified = false

	id, err := a.db.AddUser(data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	log.Println("created user", id)

	token, err := a.signer.SignAccess(&signer.AccessClaim{
		XID:    a.flakeNode.Generate().Int64(),
		UserId: id,
	})

	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})

}

func (a *Server) signUpInvite(ctx *gin.Context) {

	claim, err := a.withAccess(ctx)
	if err != nil {
		return
	}

	fmt.Println(claim)

}

type loginDetails struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}

func (a *Server) login(ctx *gin.Context) {
	data := &loginDetails{}
	err := ctx.Bind(data)
	if err != nil {
		httpx.WriteAuthErr(ctx, err)
		return
	}

	user, err := a.db.GetUserByEmail(data.Email)
	if err != nil {
		httpx.WriteAuthErr(ctx, err)
		return
	}

	// fixme => hash it
	if user.Password != data.Password {
		httpx.WriteAuthErr(ctx, err)
		return
	}

	token, err := a.signer.SignAccess(&signer.AccessClaim{
		XID:    a.flakeNode.Generate().Int64(),
		UserId: int64(user.ID),
	})

	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})

}

func (a *Server) accessMiddleware(fn func(claim *signer.AccessClaim, ctx *gin.Context) (any, error)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		claim, err := a.withAccess(ctx)
		if err != nil {
			return
		}

		resp, err := fn(claim, ctx)
		httpx.WriteJSON(ctx, resp, err)
	}

}

func (a *Server) withAccess(ctx *gin.Context) (*signer.AccessClaim, error) {

	tok := ctx.GetHeader("Authorization")
	claim, err := a.signer.ParseAccess(tok)
	if err != nil {
		httpx.WriteAuthErr(ctx, err)
		return nil, err
	}

	return claim, nil

}