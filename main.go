package main

import (
	"github.com/gin-contrib/cors"
	"log"
	"net/http"

	"github.com/arsalanaa44/vote-api/controllers"
	"github.com/arsalanaa44/vote-api/initializers"
	"github.com/arsalanaa44/vote-api/routes"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	PollController      controllers.PollController
	PollRouteController routes.PollRouteController

	VoteController      controllers.VoteController
	VoteRouteController routes.VoteRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	PollController = controllers.NewPollController(initializers.DB)
	PollRouteController = routes.NewRoutePollController(PollController)

	VoteController = controllers.NewVoteController(initializers.DB)
	VoteRouteController = routes.NewRouteVoteController(VoteController)

	server = gin.Default()
}

func main() {

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:" + config.ServerPort, config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Voting system"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	PollRouteController.PollRoute(router)
	VoteRouteController.VoteRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
