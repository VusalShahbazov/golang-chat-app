package handlers

import (
	"back/internal/helpers"
	"back/internal/models"
	"back/internal/requests"
	"back/internal/response"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func GetMessages(context *gin.Context)  {
	var json requests.GetMessages

	err := context.ShouldBindQuery(&json)
	if err != nil {
		response.GenerateValidationResponse(context, err)
		return
	}

	id := context.Param("id")
	var room models.Room
	err = models.DB.Where("id = ? " , id).First(&room).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(context)
			return
		}
		response.UnexpectedError(context , err.Error())
		return
	}

	var messages []models.Message

	err = models.DB.Preload("FromUser").Where("room_id = ? " , room.Id).
		Limit(json.Limit()).
		Offset(json.Offset()).
		Order("id desc").
		Find(&messages).Error
	if err != nil {
		response.UnexpectedError(context , err.Error())
		return
	}

	response.SuccessResponse(context , messages)
}


func SendMessage(context *gin.Context)  {
	var json requests.SendMessage

	err := context.ShouldBindJSON(&json)
	if err != nil {
		response.GenerateValidationResponse(context, err)
		return
	}
	id := context.Param("id")
	var room models.Room
	err = models.DB.Where("id = ? " , id).First(&room).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(context)
			return
		}
		response.UnexpectedError(context , err.Error())
		return
	}
	user , _ := helpers.GetUser(context)

	roomId, err  := strconv.Atoi(id)
	if err != nil {
		response.UnexpectedError(context , err.Error())
		return
	}


	message := models.Message{
		Body:       json.Body,
		FromUserId: user.Id,
		RoomId: uint64(roomId),
	}

	err = models.DB.Create(&message).Error
	if err != nil {
		response.UnexpectedError(context , err.Error())
		return
	}

	response.SuccessMessage(context)
}