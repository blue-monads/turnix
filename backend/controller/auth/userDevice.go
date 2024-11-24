package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/blue-monads/turnix/backend/services/signer"
	"github.com/blue-monads/turnix/backend/xtypes/models"
)

func (a *AuthController) ListSelfDevices(userId int64) ([]models.UserDevice, error) {
	return a.db.ListDevices(userId)
}

func (a *AuthController) AddPairDevice(userId, pid int64, name string) (string, error) {
	t := time.Now().Add(time.Hour * 24 * 7 * 3)

	token, err := a.signer.SignAccess(&signer.AccessClaim{
		UserId:          userId,
		XID:             "",
		Typeid:          signer.TokenTypePair,
		PinnedProjectId: pid,
	})
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256([]byte(token))
	hashStr := base64.URLEncoding.EncodeToString(hash[:])

	device := &models.UserDevice{
		Name:      name,
		Dtype:     "pair",
		TokenHash: hashStr,
		UserId:    userId,
		ProjectId: 0,
		ExpiresOn: &t,
	}

	did, err := a.db.AddDevice(device)
	if err != nil {
		return "", err
	}

	fullkey := fmt.Sprintf("%s--%d", token, did)

	return fullkey, nil

}

func (a *AuthController) GetSelfDevice(userId int64, id int64) (*models.UserDevice, error) {

	device, err := a.db.GetDevice(id)
	if err != nil {
		return nil, err
	}

	if device.UserId != userId {
		return nil, err
	}

	return device, nil
}

func (a *AuthController) DeleteSelfDevice(userId int64, id int64) error {

	device, err := a.db.GetDevice(id)
	if err != nil {
		return err
	}

	if device.UserId != userId {
		return err
	}

	return a.db.DeleteDevice(id)
}
