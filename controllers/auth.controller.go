package controllers

import (
	_ "context"
	"fmt"
	"log"

	// "log"
	"net/http"
	"strings"

	"aesthetic/config"
	"aesthetic/models"
	"aesthetic/services"
	"aesthetic/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthController(authService services.AuthService, userService services.UserService) *AuthController {
	return &AuthController{authService, userService}
}

func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var user models.SignUpInput

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// log.Println(&user)
	// log.Println(user)

	if user.Password != user.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	if user.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "User input empty"})
		return
	}
	
	if ac.authService == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "authService not initialized"})
		return
	}
	
	newUser, err := ac.authService.SignUpUser(&user)

	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "email already exist") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// message := "We sent an email with a verification code to " + user.Email
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "Data": &newUser})
}

func (ac *AuthController) SignInUser(ctx *gin.Context) {
	var credentials *models.SignInInput

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
	}
		
	if ac.userService == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "userService not initialized"})
		return
	}

	user, err := ac.userService.FindUserByEmail(credentials.Email)
	if err != nil {
			if err == gorm.ErrRecordNotFound {
					ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
					return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
	}

	// log.Println(credentials.Password)
	// log.Println(user.Password)

	if err := utils.VerifyPassword(user.Password, credentials.Password); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
			return
	}

	config, _ := config.LoadConfig(".")

	// Generate Tokens
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.UserID, config.AccessTokenPrivateKey)
	if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.UserID, config.RefreshTokenPrivateKey)
	if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}


func (ac *AuthController) RefreshAccessToken(ctx *gin.Context) {
	message := "could not refresh access token"

	cookie, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	config, _ := config.LoadConfig(".")

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	user, err := ac.userService.FindUserById(fmt.Sprint(sub))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.UserID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}

func (ac *AuthController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
