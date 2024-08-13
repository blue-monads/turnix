package auth

import "github.com/bornjre/turnix/backend/services/database"

type AuthController struct {
	db *database.DB
}

func NewAuthController(db *database.DB) *AuthController {
	return &AuthController{
		db: db,
	}
}
