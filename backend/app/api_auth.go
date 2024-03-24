package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bornjre/trunis/backend/database"
	"github.com/bornjre/trunis/backend/token"
	_ "github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
)

func (a *App) signUpDirect(ctx *gin.Context) {

	data := &database.User{}
	err := ctx.BindJSON(data)
	if err != nil {
		WriteErr(ctx, err)
		return
	}

	if data.Email == "" || data.Name == "" || data.Password == "" {
		return
	}

	data.IsEmailVerified = false

	id, err := a.db.AddUser(data)
	if err != nil {
		WriteErr(ctx, err)
		return
	}

	log.Println("created user", id)

	token, err := a.signer.SignAccess(&token.AccessClaim{
		XID:    a.flakeNode.Generate().Int64(),
		UserId: id,
	})

	if err != nil {
		WriteErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})

}

func (a *App) signUpInvite(ctx *gin.Context) {

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

func (a *App) login(ctx *gin.Context) {
	data := &loginDetails{}
	err := ctx.Bind(data)
	if err != nil {
		WriteAuthErr(ctx, err)
		return
	}

	user, err := a.db.GetUserByEmail(data.Email)
	if err != nil {
		WriteAuthErr(ctx, err)
		return
	}

	// fixme => hash it
	if user.Password != data.Password {
		WriteAuthErr(ctx, err)
		return
	}

	token, err := a.signer.SignAccess(&token.AccessClaim{
		XID:    a.flakeNode.Generate().Int64(),
		UserId: int64(user.ID),
	})

	if err != nil {
		WriteErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})

}

func (a *App) accessMiddleware(fn func(claim *token.AccessClaim, ctx *gin.Context) (any, error)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		claim, err := a.withAccess(ctx)
		if err != nil {
			return
		}

		resp, err := fn(claim, ctx)
		WriteJSON(ctx, resp, err)
	}

}

func (a *App) withAccess(ctx *gin.Context) (*token.AccessClaim, error) {

	tok := ctx.GetHeader("Authorization")
	claim, err := a.signer.ParseAccess(tok)
	if err != nil {
		WriteAuthErr(ctx, err)
		return nil, err
	}

	return claim, nil

}
