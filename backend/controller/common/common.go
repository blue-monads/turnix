package common

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bornjre/turnix/backend/services/database"
	"github.com/gin-gonic/gin"
	"github.com/hako/branca"
)

type CommonController struct {
	db *database.DB
	b  *branca.Branca
}

func New(db *database.DB) *CommonController {
	masterKey := make([]byte, 32)
	_, _ = rand.Read(masterKey)

	b := branca.NewBranca(string(masterKey))

	return &CommonController{
		db: db,
		b:  b,
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

func (c *CommonController) GetSharedFile(file string, ctx *gin.Context) error {
	return c.db.GetSharedFile(file, ctx.Writer)
}

type ShortFileKey struct {
	FileId    int64 `json:"f"`
	ExpiresAt int64 `json:"e"`
	UserId    int64 `json:"u"`
}

func (a *CommonController) GetFileShortKey(userId int64, id int64) (string, error) {

	file, err := a.db.GetFileMeta(id)
	if err != nil {
		return "", err
	}

	if file.OwnerUser != userId {
		scope, err := a.db.GetProjectUserScope(userId, file.OwnerProj)
		if err != nil {
			return "", err
		}

		if scope == "" {
			return "", fmt.Errorf("file not found")
		}
	}

	key := ShortFileKey{
		FileId:    file.ID,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		UserId:    userId,
	}

	out, err := json.Marshal(key)
	if err != nil {
		return "", err
	}

	return a.b.EncodeToString(string(out))

}

func (a *CommonController) GetFileWithShortKey(ctx *gin.Context) error {

	key := ctx.Param("shortkey")

	key, err := a.b.DecodeToString(key)
	if err != nil {
		return err
	}

	skey := &ShortFileKey{}
	err = json.Unmarshal([]byte(key), skey)
	if err != nil {
		return err
	}

	if skey.ExpiresAt < time.Now().Unix() {
		return fmt.Errorf("key expired")
	}

	return a.db.GetFileBlobStreaming(skey.FileId, ctx.Writer)

}
