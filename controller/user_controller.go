package controller

import (
	"kanban-golang/helper"
	"kanban-golang/model/input"
	"kanban-golang/model/response"
	"kanban-golang/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input input.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		error_message := gin.H{
			"error": helper.FormatValidationError(err),
		}

		resp := helper.APIResponse("error", error_message)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	newUser, err := h.userService.CreateUser(input)

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helper.APIResponse("error", error_message)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// success created user
	userResponse := response.UserRegisterResponse{
		ID:        newUser.ID,
		FullName:  newUser.FullName,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
	}

	resp := helper.APIResponse("success", userResponse)
	c.JSON(http.StatusCreated, resp)

}
