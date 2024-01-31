package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/arsalanaa44/vote-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VoteController struct {
	DB *gorm.DB
}

func NewVoteController(DB *gorm.DB) VoteController {
	return VoteController{DB}
}

func (bc *VoteController) CreateVote(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.VoteRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var poll models.Poll
	if result := bc.DB.Where("id = ?", payload.Poll).First(&poll); result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "poll not exist"})
		return
	}

	now := time.Now()
	newVote := models.Vote{
		User:      currentUser.ID,
		Poll:      payload.Poll,
		Choice:    payload.Choice,
		CreatedAt: now,
	}

	counts := make([]string, len(poll.Counts))
	for i := range counts {
		counts[i] = poll.Counts[i]
		if i == payload.Choice-1 {
			counter, _ := strconv.Atoi(counts[i])
			counter = counter + 1
			counts[i] = strconv.Itoa(counter)
		}
	}
	pollToUpdate := models.Poll{
		Counts: counts,
	}
	bc.DB.Model(&poll).Updates(pollToUpdate)

	result := bc.DB.Create(&newVote)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newVote.Choice})
}
func (bc *VoteController) FindVoteByPollId(ctx *gin.Context) {
	pollId := ctx.Param("pollId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var vote models.Vote
	result := bc.DB.Where(`poll = ? AND "user" = ?`, pollId, currentUser.ID).First(&vote)
	if result.Error != nil {
		vote.Choice = -1
	} else if vote.User != currentUser.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": vote.Choice})
}
