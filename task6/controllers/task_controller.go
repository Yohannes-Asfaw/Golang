package controllers

import (
	"net/http"
	"task6/data"
	"task6/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key")

func CreateUser(ctx *gin.Context) {
	var newUser models.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	errors := models.ValidateUser(newUser)
	if len(errors) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": errors})
		return
	}

	_, err := data.GetUserByUsername(ctx, newUser.Username)
	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	if err := data.CreateUser(ctx, &newUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func LoginUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	errors := models.ValidateUser(user)
	if len(errors) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": errors})
		return
	}

	retrievedUser, err := data.GetUserByUsername(ctx, user.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(retrievedUser.Password), []byte(user.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  retrievedUser.UserID,
		"username": retrievedUser.Username,
		"role":     retrievedUser.Role,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": jwtToken})
}

func Promote(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.PromoteUser(request.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user promoted to admin"})
}

func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
		return
	}

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user ID not found in token"})
		return
	}

	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	errors := models.ValidateTask(newTask)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": errors})
		return
	}

	taskID := primitive.NewObjectID()
	newTask.TaskID = taskID.Hex()
	newTask.UserID = userID.(string)

	err := data.CreateTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

func UpdateTask(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
		return
	}

	id := c.Param("id")
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	task, err := data.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
		return
	}

	id := c.Param("id")
	err := data.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
