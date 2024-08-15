package Controllers

import (
	"net/http"
	"task7/Domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase Domain.UserUseCase
}

func NewUserController(userUseCase Domain.UserUseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (uc *UserController) Register(ctx *gin.Context) {
	var newUser Domain.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := uc.UserUseCase.Register(ctx, &newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var user Domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := uc.UserUseCase.Login(ctx, user.Username, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}

func (uc *UserController) Promote(ctx *gin.Context) {
	var request struct {
		Username string `json:"username"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.UserUseCase.PromoteUser(ctx, request.Username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user promoted to admin"})
}
