package routes

import (
	"aesthetic/controllers"
	"aesthetic/middleware"
	"aesthetic/services"

	"github.com/gin-gonic/gin"
)

type ItemRouteController struct {
	itemController *controllers.ItemController
}

func NewItemRouteController(itemController *controllers.ItemController) *ItemRouteController {
	return &ItemRouteController{itemController}
}

func (ic *ItemRouteController) ItemRoute(rg *gin.RouterGroup, itemService services.ItemService, userService services.UserService) {

	router := rg.Group("items")
	router.Use(middleware.DeserializeUser(userService))
	router.GET("", ic.itemController.GetAllItem)
	router.POST("", ic.itemController.CreateItem)
	// router.PUT("/:id", ic.itemController.UpdateItem)
	// router.DELETE("/:id", ic.itemController.DeleteItem)
}
