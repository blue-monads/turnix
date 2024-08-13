package self

import "github.com/bornjre/turnix/backend/services/database"

type SelfController struct {
	db *database.DB
}

func NewSelfController(db *database.DB) *SelfController {
	return &SelfController{
		db: db,
	}
}
