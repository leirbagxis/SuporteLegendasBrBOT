package services

import (
	"telegram_bot_project/internal/api/models"

	"github.com/amarnathcjd/gogram/telegram"
)

func EnterInvitationLink(client *telegram.Client, data *models.InvitationRequest) (bool, error) {
	_, err := client.JoinChannel(data.InvitationLink)
	if err != nil {
		return false, err
	}
	return true, nil
}
