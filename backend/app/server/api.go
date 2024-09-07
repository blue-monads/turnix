package server

import (
	"strconv"

	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/gin-gonic/gin"
)

func (a *Server) userProfile(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	uid, _ := strconv.ParseInt(ctx.Param("uid"), 10, 64)
	return a.cCommon.GetUserInfo(uid)
}
