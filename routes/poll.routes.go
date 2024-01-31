package routes

import (
	"github.com/arsalanaa44/vote-api/controllers"
	"github.com/arsalanaa44/vote-api/middleware"
	"github.com/gin-gonic/gin"
)

type PollRouteController struct {
	pollController controllers.PollController
}

func NewRoutePollController(pollController controllers.PollController) PollRouteController {
	return PollRouteController{pollController}
}

func (bc *PollRouteController) PollRoute(rg *gin.RouterGroup) {

	router := rg.Group("polls")
	router.Use(middleware.DeserializeUser())
	router.POST("/", bc.pollController.CreatePoll)
	router.GET("/", bc.pollController.Findpolls)
	router.GET("/:pollId", bc.pollController.FindPollById)
}
