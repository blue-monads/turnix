package server

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/gin-gonic/gin"
)

func (a *Server) userProfile(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	uid, _ := strconv.ParseInt(ctx.Param("uid"), 10, 64)
	return a.cCommon.GetUserInfo(uid)
}

func (a *Server) getSharedFile(ctx *gin.Context) {

	file := ctx.Param("file")

	if strings.Contains(file, ".") {
		fileParts := strings.Split(file, ".")
		file = strings.Join(fileParts[:len(fileParts)-1], ".")
	}

	err := a.cCommon.GetSharedFile(file, ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

}

func (s *Server) GetFileShortKey(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	fid, err := strconv.ParseInt(ctx.Param("fid"), 10, 64)
	if err != nil {
		return nil, err
	}

	return s.cCommon.GetFileShortKey(claim.UserId, fid)
}

func (s *Server) GetFileWithShortKey(ctx *gin.Context) {
	err := s.cCommon.GetFileWithShortKey(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

}
