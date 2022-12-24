package activity

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-list/database"
	"todo-list/models"

	"github.com/gin-gonic/gin"
)

type emptyStruct struct{}

var emptyData emptyStruct

const (
	successMessage = "Success"
	notFound       = "Not Found"
	badRequest     = "Bad Request"
	cannotNull     = "cannot be null"
)

func CreateActivity(ctx *gin.Context) {
	db := database.GetDB()
	activity := models.Activity{}

	err := ctx.ShouldBindJSON(&activity)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if activity.Title == "" {
		ctx.JSON(http.StatusBadRequest, resultResponse{
			badRequest,
			fmt.Sprintf("title %s", cannotNull),
			emptyData,
		})
		return
	}

	err = db.Create(&activity).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resultResponse{
			badRequest, err.Error(), emptyData,
		})
	}

	ctx.JSON(http.StatusCreated, resultResponse{
		successMessage, successMessage, createActivity{
			ID:        activity.ID,
			Email:     activity.Email,
			Title:     activity.Title,
			CreatedAt: activity.CreatedAt,
			UpdatedAt: activity.UpdatedAt,
		},
	})
}

func GetActivityByID(ctx *gin.Context) {
	db := database.GetDB()
	activity := models.Activity{}
	activityID := ctx.Param("id")

	err := db.First(&activity, activityID).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, resultResponse{
			notFound,
			fmt.Sprintf("Activity with ID %s Not Found", activityID),
			emptyData,
		})
		return
	}

	ctx.JSON(http.StatusOK, resultResponse{
		successMessage, successMessage, activity,
	})
}

func GetAllActivities(ctx *gin.Context) {
	db := database.GetDB()
	activities := []models.Activity{}

	err := db.Find(&activities).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resultResponse{
			badRequest, err.Error(), emptyData,
		})
		return
	}
	ctx.JSON(http.StatusOK, resultResponse{
		successMessage, successMessage, activities,
	})
}

func UpdateActivity(ctx *gin.Context) {
	db := database.GetDB()
	activity := models.Activity{}
	requestBody := models.Activity{}
	activityID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	activity.ID = uint(activityID)
	err = db.First(&activity).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, resultResponse{
			notFound,
			fmt.Sprintf("Activity with ID %v Not Found", activityID),
			emptyData,
		})
		return
	}

	if requestBody.Title == "" {
		ctx.JSON(http.StatusBadRequest, resultResponse{
			badRequest,
			fmt.Sprintf("title %s", cannotNull),
			emptyData,
		})
		return
	}

	if requestBody.Email == "" {
		requestBody.Email = activity.Email
	}

	activity.Title = requestBody.Title
	activity.Email = requestBody.Email

	err = db.Save(&activity).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resultResponse{
		successMessage, successMessage, activity,
	})
}

func DeleteActivity(ctx *gin.Context) {
	db := database.GetDB()
	activity := models.Activity{}
	activityID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	activity.ID = uint(activityID)

	err = db.First(&activity).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, resultResponse{
			notFound,
			fmt.Sprintf("Activity with ID %v Not Found", activityID),
			emptyData,
		})
		return
	}

	err = db.Delete(&activity).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resultResponse{
		successMessage, successMessage, emptyData,
	})
}
