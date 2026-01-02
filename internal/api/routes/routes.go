package routes

import (
	"telegram_bot_project/internal/api/controllers"

	"github.com/amarnathcjd/gogram/telegram"
	"github.com/gin-gonic/gin"
)

func SetupRouter(client *telegram.Client) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/status", controllers.GetStatus)
		api.GET("/bot/info", controllers.GetBotInfo)

		api.PUT("/bot/caption", controllers.CaptionController(client))
		api.POST("/bot/enterinvitation", controllers.EnterInvitationController(client))
	}

	return r
}
