package todo

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-list/database"
	"todo-list/models"

	"github.com/gin-gonic/gin"
)

const (
	successMessage = "Success"
	emptyData      = "{}"
	notFound       = "Not Found"
	badRequest     = "Bad Request"
	cannotNull     = "cannot be null"
)

func CreateTodo(ctx *gin.Context) {
	db := database.GetDB()
	todo := models.ToDo{}

	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if todo.Title == "" {
		ctx.JSON(http.StatusBadRequest, resultResponse{
			badRequest,
			fmt.Sprintf("title %s", cannotNull),
			emptyData,
		})
		return
	}

	if todo.ActivityID == 0 {
		ctx.JSON(http.StatusBadRequest, resultResponse{
			badRequest,
			fmt.Sprintf("activity_group_id %s", cannotNull),
			emptyData,
		})
		return
	}

	err = db.Create(&todo).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resultResponse{
			badRequest, err.Error(), emptyData,
		})
	}

	ctx.JSON(http.StatusCreated, resultResponse{
		successMessage, successMessage, createTodo{
			CreatedAt:  todo.CreatedAt,
			UpdatedAt:  todo.UpdatedAt,
			ID:         todo.ID,
			Title:      todo.Title,
			ActivityID: todo.ActivityID,
			IsActive:   todo.IsActive,
			Priority:   todo.Priority,
		},
	})
}

func GetTodoByID(ctx *gin.Context) {
	db := database.GetDB()
	todo := models.ToDo{}
	todoID := ctx.Param("id")

	err := db.First(&todo, todoID).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, resultResponse{
			notFound,
			fmt.Sprintf("Todo with ID %s Not Found", todoID),
			emptyData,
		})
		return
	}

	ctx.JSON(http.StatusOK, resultResponse{
		successMessage, successMessage, todo,
	})
}

func GetAllTodo(ctx *gin.Context) {
	db := database.GetDB()
	todo := []models.ToDo{}

	err := db.Find(&todo).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, resultResponse{
			badRequest, err.Error(), emptyData,
		})
		return
	}

	ctx.JSON(http.StatusOK, resultResponse{
		successMessage, successMessage, todo,
	})
}

func UpdateTodo(ctx *gin.Context) {
	db := database.GetDB()
	todo := models.ToDo{}
	requestBody := models.ToDo{}
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	todo.ID = uint(todoID)
	err = db.First(&todo).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, resultResponse{
			notFound,
			fmt.Sprintf("Todo with ID %v Not Found", todoID),
			emptyData,
		})
		return
	}

	if requestBody.Title == "" {
		requestBody.Title = todo.Title
	}

	if !requestBody.IsActive {
		requestBody.IsActive = todo.IsActive
	}

	if requestBody.Priority == "" {
		requestBody.Priority = todo.Priority
	}

	todo.Title = requestBody.Title
	todo.IsActive = requestBody.IsActive
	todo.Priority = requestBody.Priority

	err = db.Save(&todo).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resultResponse{
		successMessage, successMessage, todo,
	})
}

func DeleteTodo(ctx *gin.Context) {
	db := database.GetDB()
	todo := models.ToDo{}
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	todo.ID = uint(todoID)

	err = db.First(&todo).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, resultResponse{
			notFound,
			fmt.Sprintf("Todo with ID %v Not Found", todoID),
			emptyData,
		})
		return
	}

	err = db.Delete(&todo).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resultResponse{
		successMessage, successMessage, emptyData,
	})
}
