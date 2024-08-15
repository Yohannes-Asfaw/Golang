package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"task7/Delivery/Controllers"
	"task7/Infrastructure"
	"task7/Usecases"
	"task7/Repositories"
)

func SetupRouter(db *mongo.Database) *gin.Engine {
	// Initialize repositories
	userRepo := Repositories.NewUserRepository(db.Collection("users"))
	taskRepo := Repositories.NewTaskRepository(db.Collection("tasks"))

	jwtService := Infrastructure.NewJWTService("your_secret_key", "your_issuer")
	passwordService := Infrastructure.NewPasswordService()
	userUseCase := Usecases.NewUserUseCase(userRepo, jwtService, passwordService)

	authService := Infrastructure.NewAuthenticationService("your_secret_key")
	authzService := Infrastructure.NewAuthorizationService()

	taskUseCase := Usecases.NewTaskUseCase(taskRepo)

	userController := Controllers.NewUserController(userUseCase)
	taskController := Controllers.NewTaskController(taskUseCase)

	// Set up the Gin router
	r := gin.Default()

	// Public routes
	r.POST("/login", userController.Login)
	r.POST("/register", userController.Register)

	// Protected task routes
	taskRoutes := r.Group("/tasks")
	taskRoutes.Use(Infrastructure.AuthMiddleware(authService))
	{
		taskRoutes.GET("/", taskController.GetTasks)
		taskRoutes.GET("/:id", taskController.GetTaskByID)
	}

	taskRoutes.Use(Infrastructure.AuthMiddleware(authService), Infrastructure.AdminOnlyMiddleware(authzService))
	{
		taskRoutes.POST("/", taskController.CreateTask)
		taskRoutes.PUT("/:id", taskController.UpdateTask)
		taskRoutes.DELETE("/:id", taskController.DeleteTask)
	}

	return r
}
