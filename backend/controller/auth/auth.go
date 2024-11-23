package auth

import (
	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/services/signer"
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
