package routes

import (
	controller "github.com/novitaekasari/restraurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controller.UpdateMenu())
}