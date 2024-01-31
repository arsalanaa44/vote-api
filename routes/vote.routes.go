package routes

import (
	"github.com/arsalanaa44/vote-api/controllers"
	"github.com/arsalanaa44/vote-api/middleware"
	"github.com/gin-gonic/gin"
)

type VoteRouteController struct {
	voteController controllers.VoteController
}

func NewRouteVoteController(voteController controllers.VoteController) VoteRouteController {
	return VoteRouteController{voteController}
}

func (bc *VoteRouteController) VoteRoute(rg *gin.RouterGroup) {

	router := rg.Group("votes")
	router.Use(middleware.DeserializeUser())
	router.POST("/", bc.voteController.CreateVote)
	router.GET("/:pollId", bc.voteController.FindVoteByPollId)
}
