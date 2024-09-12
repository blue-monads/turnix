package server

import (
	"net/http"
	"strconv"

	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/gin-gonic/gin"
)

func (a *Server) userProfile(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	uid, _ := strconv.ParseInt(ctx.Param("uid"), 10, 64)
	return a.cCommon.GetUserInfo(uid)
}

func (a *Server) getSharedFile(ctx *gin.Context) {
	bytes, err := a.cCommon.GetSharedFile(ctx.Param("file"))
	if err != nil {
		return
	}

	ctx.Writer.Write(bytes)
	ctx.Writer.Header().Set("Content-Type", "application/octet-stream")
	ctx.Writer.WriteHeader(http.StatusOK)

}
