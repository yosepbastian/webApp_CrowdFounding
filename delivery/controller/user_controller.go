package controller

import (
	"net/http"
	"web-app-crowdfounding/helper"
	"web-app-crowdfounding/models"
	"web-app-crowdfounding/usecase"

	"github.com/gin-gonic/gin"
)

type userController struct {
	router      *gin.RouterGroup
	userUseCase usecase.UserUseCase
}

func (u *userController) GetAllUser(ctx *gin.Context) {
	user, err := u.userUseCase.GetAllUser()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"user": user,
	})

}

func (u *userController) RegisterUser(ctx *gin.Context) {

	var userInput models.RegisterUserInput

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		formatErr := helper.FormatValidationErr(err)
		errorMessage := gin.H{"errors": formatErr}
		response := helper.ApiResponse("Register account Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	user, err := u.userUseCase.RegisterUser(userInput)

	if err != nil {
		formatErr := helper.FormatValidationErr(err)
		errorMessage := gin.H{"errors": formatErr}
		response := helper.ApiResponse("Register account Failed", http.StatusBadRequest, "Error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helper.FormatUser(user, "")
	response := helper.ApiResponse("Account has been Registered", http.StatusCreated, "success", formatter)
	ctx.JSON(http.StatusCreated, response)

}

func (u *userController) Login(ctx *gin.Context) {
	var userInput models.LoginUser

	err := ctx.ShouldBindJSON(&userInput)
	if err != nil {
		formatErr := helper.FormatValidationErr(err)
		errorMessage := gin.H{"errors": formatErr}
		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	loginuser, err := u.userUseCase.Login(userInput)
	if err != nil {

		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := helper.FormatUser(loginuser, "")
	response := helper.ApiResponse("Login success", http.StatusOK, "success", formatter)
	ctx.JSON(http.StatusOK, response)

}

func NewUserController(routerGroup *gin.RouterGroup, userUC usecase.UserUseCase) *userController {
	newUserController := userController{
		router:      routerGroup,
		userUseCase: userUC,
	}

	newUserController.router.POST("/users", newUserController.RegisterUser)
	newUserController.router.POST("/sessions", newUserController.Login)
	return &newUserController
}
