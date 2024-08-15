package Controllers

import (
	"context"
	"net/http"
	"task7/Domain"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUseCase Domain.TaskUseCase
}

func NewTaskController(taskUseCase Domain.TaskUseCase) *TaskController {
	return &TaskController{
		TaskUseCase: taskUseCase,
	}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var newTask Domain.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	createdTask, err := tc.TaskUseCase.Create(context.Background(), &newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	var updatedTask Domain.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	task, err := tc.TaskUseCase.Update(context.Background(), taskID, &updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	err := tc.TaskUseCase.Delete(context.Background(), taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.TaskUseCase.GetAll(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTaskByID(c *gin.Context) {
	taskID := c.Param("id")
	task, err := tc.TaskUseCase.GetById(context.Background(), taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}
