package routers

import (
	"todo-list/controllers/activity"
	"todo-list/controllers/todo"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	activityRouter := r.Group("/activities")
	{
		activityRouter.POST("/", activity.CreateActivity)
		activityRouter.GET("/", activity.GetAllActivities)
		activityRouter.GET("/:id", activity.GetActivityByID)
		activityRouter.PATCH("/:id", activity.UpdateActivity)
		activityRouter.DELETE("/:id", activity.DeleteActivity)
	}

	todoRouter := r.Group("/todos")
	{
		todoRouter.POST("/", todo.CreateTodo)
		todoRouter.GET("/", todo.GetAllTodo)
		todoRouter.GET("/:id", todo.GetTodoByID)
		todoRouter.PATCH("/:id", todo.UpdateTodo)
		todoRouter.DELETE("/:id", todo.DeleteTodo)
	}

	return r
}
