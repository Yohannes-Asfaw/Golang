package controllers

import (
    "net/http"
    "task4/data"
    "task4/models"
    "github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
    tasks := data.GetTasks()
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
    var newTask models.Task
    if err := c.ShouldBindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    data.CreateTask(newTask)
    c.JSON(http.StatusCreated, newTask)
}

func UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var updatedTask models.Task
    if err := c.ShouldBindJSON(&updatedTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := data.UpdateTask(id, updatedTask)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
    id := c.Param("id")
    err := data.DeleteTask(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
