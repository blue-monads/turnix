package books

import (
	"github.com/bornjre/trunis/backend/xtypes"
	"github.com/gin-gonic/gin"
)

func (b *BookModule) register(group *gin.RouterGroup) error {

	x := b.app.AuthMiddleware

	accGrp := group.Group("/:pid/account")

	accGrp.GET("/", x(b.listAccount))
	accGrp.POST("/", x(b.addAccount))
	accGrp.GET("/:id", x(b.getAccount))
	accGrp.POST("/:id", x(b.updateAccount))
	accGrp.DELETE("/:id", x(b.deleteAccount))

	//	txnGroup := group.Group("/:pid/account")

	return nil
}

func (b *BookModule) listAccount(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	return b.dbOpListAccount(pid, ctx.Claim.UserId)
}

func (b *BookModule) addAccount(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &Account{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return b.dbOpAddAccount(pid, ctx.Claim.UserId, data)
}

func (b *BookModule) getAccount(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	return b.dbOpGetAccount(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
}

func (b *BookModule) updateAccount(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpUpdateAccount(pid, ctx.Claim.UserId, ctx.ParamInt64("id"), data)
	if err != nil {
		return nil, err
	}

	return nil, nil

}

func (b *BookModule) deleteAccount(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := make(map[string]any)
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	err = b.dbOpDeleteAccount(pid, ctx.Claim.UserId, ctx.ParamInt64("id"))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
