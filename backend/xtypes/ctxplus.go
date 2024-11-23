package xtypes

import (
	"strconv"

	"github.com/blue-monads/turnix/backend/services/signer"
	"github.com/gin-gonic/gin"
)

type ContextPlus struct {
	Claim *signer.AccessClaim
	Http  *gin.Context
}

func (c *ContextPlus) ProjectId() int64 {

	id, err := strconv.ParseInt(c.Http.Param("pid"), 10, 64)
	if err != nil {
		panic(err)
	}

	return id
}

func (c *ContextPlus) ParamInt64(name string) int64 {
	id, err := strconv.ParseInt(c.Http.Param(name), 10, 64)
	if err != nil {
		panic(err)
	}

	return id
}
