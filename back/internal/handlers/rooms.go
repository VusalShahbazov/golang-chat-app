package handlers

import (
	"back/internal/helpers"
	"back/internal/models"
	"back/internal/requests"
	"back/internal/response"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRooms(context *gin.Context)  {
	var rooms []models.Room
	models.DB.Preload("Owner").
		Order("id desc").
		Find(&rooms)
	response.SuccessResponse(context , rooms)
}

func GetRoom(context *gin.Context)  {
	id := context.Param("id")
	var room models.Room
	models.DB.Preload("Owner").Where("id = ? " , id).First(&room)
	response.SuccessResponse(context , room)
}

func UpdateRoom(context *gin.Context)  {
	var json requests.StoreRoom

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

	if room.OwnerId != user.Id {
		response.Forbidden(context , gin.H{"error" : "you are not owner"})
		return
	}

	room.Name = json.Name

	err = models.DB.Save(&room).Error
	if err != nil {
		response.UnexpectedError(context , err.Error())
		return
	}

	response.SuccessResponse(context , room)
}
func StoreRoom(context *gin.Context)  {
	var json requests.StoreRoom

	err := context.ShouldBindJSON(&json)
	if err != nil {
		response.GenerateValidationResponse(context, err)
		return
	}
	user , _ := helpers.GetUser(context)
	room:= models.Room{
		Name:    json.Name,
		OwnerId: user.Id,
	}
	err = models.DB.Create(&room).Error
	if err != nil {
		response.UnexpectedError(context , err.Error())
		return
	}
	response.SuccessResponse(context , room)
}

func DeleteRoom(context *gin.Context)  {
	id := context.Param("id")
	var room models.Room
	err := models.DB.Where("id = ? " , id).First(&room).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(context)
			return
		}
		response.UnexpectedError(context , err.Error())
		return
	}

	user , _ := helpers.GetUser(context)

	if room.OwnerId != user.Id {
		response.Forbidden(context , gin.H{"error" : "you are not owner"})
		return
	}

	models.DB.Where("room_id = ? " ,  room.Id).Delete(&models.Message{})

	models.DB.Delete(&room)

	response.SuccessMessage(context)
}