package routes

import (
	"aesthetic/controllers"
	"aesthetic/middleware"
	"aesthetic/services"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController *controllers.UserController
}

func NewRouteUserController(userController *controllers.UserController) *UserRouteController {
	return &UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {

	router := rg.Group("users")
	router.Use(middleware.DeserializeUser(userService))
	router.GET("/me", uc.userController.GetMe)
}
