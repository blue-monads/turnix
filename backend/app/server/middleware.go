package server

import (
	"github.com/blue-monads/turnix/backend/utils/libx/httpx"
	"github.com/blue-monads/turnix/backend/xtypes"

	"github.com/gin-gonic/gin"
)

func (a *Server) AuthMiddleware(fn xtypes.ApiHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		claim, err := a.withAccess(ctx)
		if err != nil {
			return
		}

		resp, err := fn(xtypes.ContextPlus{
			Claim: claim,
			Http:  ctx,
		})

		httpx.WriteJSON(ctx, resp, err)
	}
}
