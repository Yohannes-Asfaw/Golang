package router

import (
	"task6/controllers"
	"task6/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes (No authentication required)
	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.LoginUser)

	api := r.Group("")
	api.Use(middleware.Authenticate()) // Apply authentication middleware to all routes in this group
	{
		api.GET("/tasks", controllers.GetTasks)
		api.GET("/tasks/:id", controllers.GetTaskByID)

		// Only admins can create, update, or delete tasks
		api.POST("/tasks", middleware.Authorize("admin"), controllers.CreateTask)
		api.PUT("/tasks/:id", middleware.Authorize("admin"), controllers.UpdateTask)
		api.DELETE("/tasks/:id", middleware.Authorize("admin"), controllers.DeleteTask)

		// Only admins can promote other users
		api.POST("/promote/:username", middleware.Authorize("admin"), controllers.Promote)
	}

	return r
}
