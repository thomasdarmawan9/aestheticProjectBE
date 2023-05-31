package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"aesthetic/config"
	"aesthetic/controllers"
	"aesthetic/models"
	"aesthetic/routes"
	"aesthetic/services"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	server              *gin.Engine
	ctx                 context.Context
	db                  *gorm.DB
	redisClient         *redis.Client
	userService         services.UserService
	userController      *controllers.UserController
	userRouteController *routes.UserRouteController
	authService         services.AuthService
	authController      *controllers.AuthController
	authRouteController *routes.AuthRouteController
	itemService         services.ItemService
	itemController      *controllers.ItemController
	itemRouteController *routes.ItemRouteController
)

func init() {
	var err error
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ctx = context.TODO()

	// Connect to PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.PostgresHost,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDB,
		config.PostgresPort,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("PostgreSQL connect fail")
		panic(err)
	}

	// Connect to Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr: config.RedisUri,
	})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Redis client connected successfully...")

	// Migrate the models
	err = db.Debug().AutoMigrate(&models.Tb_Customers{}, &models.Item{}, &models.Transaction{}, &models.PaymentMethodGroup{}, &models.PaymentMethodItem{})
	if err != nil {
		fmt.Println("Error migrating User model")
		panic(err)
	}

	// Initialize the services
	userService = services.NewUserService(db)
	authService = services.NewAuthService(db, ctx)
	itemService = services.NewItemService(db)

	// Initialize the controllers
	authController = controllers.NewAuthController(authService, userService)
	userController = controllers.NewUserController(userService)
	itemController = controllers.NewItemController(itemService)

	// Initialize the routes
	authRouteController = routes.NewAuthRouteController(authController)
	userRouteController = routes.NewRouteUserController(userController)
	itemRouteController = routes.NewItemRouteController(itemController)

	fmt.Println("PostgreSQL successfully connected...")
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config", err)
	}
	startGinServer(config)
}

func startGinServer(config config.Config) {
	value, err := redisClient.Get(ctx, "test").Result()

	if err == redis.Nil {
		fmt.Println("key: test does not exist")
	} else if err != nil {
		panic(err)
	}

	server = gin.Default()

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": value})
	})

	authRouteController.AuthRoute(router, userService)
	userRouteController.UserRoute(router, userService)
	itemRouteController.ItemRoute(router, itemService, userService)

	fmt.Println("routes running")

	server.Run(":" + config.Port)
}
