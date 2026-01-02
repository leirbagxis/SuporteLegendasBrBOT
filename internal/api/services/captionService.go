package services

import (
	"telegram_bot_project/internal/api/models"

	"github.com/amarnathcjd/gogram/telegram"
)

func EditCaption(client *telegram.Client, data *models.CaptionRequest) error {
	_, err := client.EditMessage(data.ChannelID, int32(data.MessageID), data.Caption)
	if err != nil {
		return err
	}
	return nil
}
