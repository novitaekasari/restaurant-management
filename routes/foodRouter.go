package routes

import (
	controller "github.com/novitaekasari/restraurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
}

// package routes

// import (
// 	controller "github.com/novitaekasari/restraurant-management/controllers"

// 	"github.com/gin-gonic/gin"
// )

// func FoodRoutes(incomingRoutes *gin.Engine) {
// 	incomingRoutes.GET("/foods", controller.GetFoods())
// 	incomingRoutes.GET("/foods/:foods_id", controller.GetFood())
// 	incomingRoutes.POST("/foods", controller.CreateFood())
// 	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
// }
