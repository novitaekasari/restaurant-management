package routes

import (
	controller "github.com/novitaekasari/restraurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("orderItems", controller.GetOrderItems())
	incomingRoutes.GET("orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("orderItems/:orderItem_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("orderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("orderItems/:orderItem_id", controller.UpdateOrderItem())
}
