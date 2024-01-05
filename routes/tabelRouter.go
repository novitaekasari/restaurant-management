package routes

import (
	controller "github.com/novitaekasari/restraurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func TabelRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("tabels", controller.GetTabels())
	incomingRoutes.GET("tabels/:order_id", controller.GetTabel())
	incomingRoutes.POST("tabels", controller.CreateTabel())
	incomingRoutes.PATCH("tabels/:order_id", controller.UpdateTabel())
}
