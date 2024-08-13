package auth

import (
	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"
)

type AuthController struct {
	db     *database.DB
	signer *signer.Signer
}

func NewAuthController(db *database.DB) *AuthController {
	return &AuthController{
		db: db,
	}
}
