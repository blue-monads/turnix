package common

import "github.com/bornjre/turnix/backend/services/database"

type CommonController struct {
	db *database.DB
}

func New(db *database.DB) *CommonController {
	return &CommonController{
		db: db,
	}
}

type UserInfo struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Utype    string  `json:"utype"`
	Email    string  `json:"email"`
	Phone    *string `json:"phone"`
	Username *string `json:"username"`
	Bio      string  `json:"bio"`
}

func (c *CommonController) GetUserInfo(uid int64) (*UserInfo, error) {

	user, err := c.db.GetUser(uid)

	if err != nil {
		return nil, err
	}

	return &UserInfo{
		ID:       user.ID,
		Name:     user.Name,
		Utype:    user.Utype,
		Email:    user.Email,
		Phone:    user.Phone,
		Username: user.Username,
		Bio:      user.Bio,
	}, nil

}

func (c *CommonController) GetSharedFile(file string) ([]byte, error) {
	return c.db.GetSharedFile(file)
}
