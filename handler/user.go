package handler

import (
	"horus/auth"
	"horus/helper"
	"horus/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//catch json input from user
	//map the input to RegisterUserInput
	//Passing input to service

	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register user failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Generate token failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatJSONUser(newUser, token)
	response := helper.APIResponse("Account has been created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
func (h *userHandler) LoginUser(c *gin.Context) {
	//catch user input by handler
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	LoggedinUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(LoggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	formatter := user.FormatJSONUser(LoggedinUser, token)
	response := helper.APIResponse("Successfully Logged in", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
	//mapping user input into struct LoginInputUser
	//passing input struct to service
	//service will match the user input in db by his email
	//matching the password after match the email is true
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")

	// Parse the string into a UUID
	id, err := uuid.Parse(idParam)
	if err != nil {
		// Handle invalid UUID error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	// Fetch the user by UUID
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		// Handle error from the service
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
