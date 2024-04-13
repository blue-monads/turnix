package app

import (
	"github.com/bornjre/trunis/backend/xtypes"
	"github.com/gin-gonic/gin"
)

func (a *App) AuthMiddleware(fn xtypes.ApiHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		claim, err := a.withAccess(ctx)
		if err != nil {
			return
		}

		resp, err := fn(claim, ctx)
		WriteJSON(ctx, resp, err)
	}
}
