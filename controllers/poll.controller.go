package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/arsalanaa44/vote-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PollController struct {
	DB *gorm.DB
}

func NewPollController(DB *gorm.DB) PollController {
	return PollController{DB}
}

func (bc *PollController) CreatePoll(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.PollRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	counts := make([]string, len(payload.Options))
	for i := range counts {
		counts[i] = "0"
	}

	now := time.Now()
	newPoll := models.Poll{
		Description: payload.Description,
		Options:     models.StringArray(payload.Options),
		User:        currentUser.ID,
		Counts:      counts,
		CreatedAt:   now,
	}

	result := bc.DB.Create(&newPoll)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPoll})
}
func (bc *PollController) FindPollById(ctx *gin.Context) {
	pollId := ctx.Param("pollId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var poll models.Poll
	result := bc.DB.First(&poll, "id = ?", pollId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No poll found"})
		return
	}
	if poll.User != currentUser.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": poll})
}

func (bc *PollController) Findpolls(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")
	currentUser := ctx.MustGet("currentUser").(models.User)

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var polls []models.Poll
	results := bc.DB.Where(`"user" = ? `, currentUser.ID).Limit(intLimit).Offset(offset).Find(&polls)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(polls), "data": polls})
}
