package controllers

import (
	"fmt"
	"net/http"
	"telegram_bot_project/internal/api/models"
	"telegram_bot_project/internal/api/services"

	"github.com/amarnathcjd/gogram/telegram"
	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	response := models.StatusResponse{
		Status:  "online",
		Message: "Bot e API estão rodando perfeitamente.",
		Version: "1.0.0-MVP",
	}
	c.JSON(http.StatusOK, response)
}

func GetBotInfo(c *gin.Context) {
	// Exemplo de retorno estático para o MVP
	info := models.BotInfo{
		Username: "ExemplarBot",
		ID:       123456789,
		IsActive: true,
	}
	c.JSON(http.StatusOK, info)
}

func CaptionController(client *telegram.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var channelData models.CaptionRequest
		if err := c.ShouldBindJSON(&channelData); err != nil {
			c.JSON(http.StatusBadRequest, models.CaptionResponse{
				Success: false,
				Message: "Dados inválidos: " + err.Error(),
			})
			return
		}

		err := services.EditCaption(client, &channelData)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.CaptionResponse{
				Success: false,
				Message: "erro ao editar mensagem " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, models.CaptionResponse{
			Success: true,
			Message: "Legenda editada com sucesso",
		})
	}
}

func EnterInvitationController(client *telegram.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var invitationData models.InvitationRequest
		if err := c.ShouldBindJSON(&invitationData); err != nil {
			c.JSON(http.StatusBadRequest, models.CaptionResponse{
				Success: false,
				Message: "Dados inválidos: " + err.Error(),
			})
			return
		}

		res, err := services.EnterInvitationLink(client, &invitationData)
		if err != nil || !res {
			c.JSON(http.StatusBadRequest, models.CaptionResponse{
				Success: false,
				Message: "erro ao entrar no canal " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, models.CaptionResponse{
			Success: true,
			Message: fmt.Sprintf("%s", res),
		})
	}
}
