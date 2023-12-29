package routes

import (
	controller "github.com/novitaekasari/restraurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func TabelRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetTabels())
	incomingRoutes.GET("/orders/:order_id", controller.GetTabels())
	incomingRoutes.POST("/orders", controller.CreateTabel())
	incomingRoutes.PATCH("/orders/:order_id", controller.UpdateTabel())
}
