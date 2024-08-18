package self

import "github.com/bornjre/turnix/backend/xtypes/models"

func (a *SelfController) ListUserMessages(userId int64, count, cursor int64) ([]models.UserMessage, error) {

	return a.db.ListUserMessages(userId, count, cursor)
}

func (a *SelfController) SendUserMessage(fromUserId int64, toUserId int64) error {

	// a.db.AddUserMessage(&models.UserMessage{
	// 	ToUser:        userId,
	// 	FromUserId:    userId,
	// 	FromProjectId: 0,
	// 	CallbackToken: "",
	// 	WarnLevel:     0,
	// 	CreatedAt:     0,
	// })

	return nil
}
