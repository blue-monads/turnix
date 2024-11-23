package self

import "github.com/blue-monads/turnix/backend/xtypes/models"

func (a *SelfController) ListUserMessages(userId int64, count, cursor int64) ([]models.UserMessage, error) {

	return a.db.ListUserMessages(userId, count, cursor)
}

type MessegeRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (a *SelfController) SendUserMessage(fromUserId int64, toUserId int64, data *MessegeRequest) error {

	_, err := a.db.AddUserMessage(&models.UserMessage{
		Name:     data.Title,
		Contents: data.Body,
		ToUser:   toUserId,
		FromUser: fromUserId,
		Type:     "user",
		IsRead:   false,
	})

	return err
}
