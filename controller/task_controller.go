package controller

import (
	"kanban-golang/helper"
	"kanban-golang/model/input"
	"kanban-golang/model/response"
	"kanban-golang/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type taskController struct {
	taskService service.TaskService
}

func NewTaskController(taskService service.TaskService) *taskController {
	return &taskController{taskService}
}

func (h *taskController) CreateNewTask(c *gin.Context) {
	var input input.TaskCreateInput

	err := c.ShouldBindJSON(&input)
	currentUser := c.MustGet("currentUser").(int)

	if err != nil {
		error_message := gin.H{
			"error": helper.FormatValidationError(err),
		}

		resp := helper.APIResponse("error", error_message)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	newTask, err := h.taskService.CreateTask(input, currentUser)

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helper.APIResponse("error", error_message)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// success create new task
	taskResponse := response.TaskCreatedResponse{
		ID:          newTask.ID,
		Title:       newTask.Title,
		Description: input.Description,
		Status:      newTask.Status,
		UserID:      newTask.UserID,
		CategoryID:  input.CategoryID,
		CreatedAt:   newTask.CreatedAt,
	}

	resp := helper.APIResponse("success", taskResponse)
	c.JSON(http.StatusCreated, resp)
}

func (h *taskController) GetAllTask(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(int)

	allTask, err := h.taskService.GetTasks(currentUser)

	if err != nil {
		error_message := gin.H{
			"error": helper.FormatValidationError(err),
		}

		resp := helper.APIResponse("error", error_message)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helper.APIResponse("success", allTask)
	c.JSON(http.StatusOK, resp)
}

func (h *taskController) UpdateTaskStatus(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(int)

	var inputTask input.TaskUpdateStatus
	var inputID input.TaskID

	err := c.ShouldBindJSON(&inputTask)
	err = c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// check first authenticate user & task
	taskWillEdit, err := h.taskService.GetTaskDetail(inputID.ID)

	if taskWillEdit.ID == 0 {
		response := helper.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusNotFound, response)
		return
	}

	if taskWillEdit.UserID != currentUser {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	// will lolos

	updatedTask, err := h.taskService.UpdateStatusTask(currentUser, inputTask)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskResponse := response.TaskUpdatedResponse{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		UserID:      updatedTask.UserID,
		CategoryID:  updatedTask.CategoryID,
		UpdatedAt:   updatedTask.UpdatedAt,
	}

	response := helper.APIResponse("ok", taskResponse)
	c.JSON(http.StatusOK, response)
}

func (h *taskController) UpdateTask(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(int)

	var inputTask input.TaskEditInput
	var inputID input.TaskID

	err := c.ShouldBindJSON(&inputTask)
	err = c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// check first authenticate user & task
	taskWillEdit, err := h.taskService.GetTaskDetail(inputID.ID)

	if taskWillEdit.ID == 0 {
		response := helper.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusNotFound, response)
		return
	}

	if taskWillEdit.UserID != currentUser {
		response := helper.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	updatedTask, err := h.taskService.UpdateTask(inputID.ID, inputTask)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskResponse := response.TaskUpdatedResponse{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		UserID:      updatedTask.UserID,
		CategoryID:  updatedTask.CategoryID,
		UpdatedAt:   updatedTask.UpdatedAt,
	}

	response := helper.APIResponse("ok", taskResponse)
	c.JSON(http.StatusOK, response)
}

func (h *taskController) UpdateStatusTask(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var inputTaskStatus input.TaskUpdateStatus
	var inputID input.TaskID

	err := c.ShouldBindJSON(&inputTaskStatus)
	err = c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskWillEdit, err := h.taskService.GetTaskDetail(inputID.ID)

	if taskWillEdit.ID == 0 {
		response := helper.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusNotFound, response)
		return
	}

	if taskWillEdit.UserID != currentUser {
		response := helper.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	updatedStatusTask, err := h.taskService.UpdateStatusTask(inputID.ID, inputTaskStatus)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskResponse := response.TaskUpdatedResponse{
		ID:          updatedStatusTask.ID,
		Title:       updatedStatusTask.Title,
		Description: updatedStatusTask.Description,
		Status:      updatedStatusTask.Status,
		UserID:      updatedStatusTask.UserID,
		CategoryID:  updatedStatusTask.CategoryID,
		UpdatedAt:   updatedStatusTask.UpdatedAt,
	}

	response := helper.APIResponse("ok", taskResponse)
	c.JSON(http.StatusOK, response)

}

func (h *taskController) UpdateCategoryTask(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var inputCategory input.TaskUpdateCategory
	var inputID input.TaskID

	err := c.ShouldBindJSON(&inputCategory)
	err = c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskWillEdit, err := h.taskService.GetTaskDetail(inputID.ID)

	if taskWillEdit.ID == 0 {
		response := helper.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusNotFound, response)
		return
	}

	if taskWillEdit.UserID != currentUser {
		response := helper.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	updatedStatusTask, err := h.taskService.UpdateStatusCategoryTask(inputID.ID, inputCategory)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskResponse := response.TaskUpdatedResponse{
		ID:          updatedStatusTask.ID,
		Title:       updatedStatusTask.Title,
		Description: updatedStatusTask.Description,
		Status:      updatedStatusTask.Status,
		UserID:      updatedStatusTask.UserID,
		CategoryID:  updatedStatusTask.CategoryID,
		UpdatedAt:   updatedStatusTask.UpdatedAt,
	}

	response := helper.APIResponse("ok", taskResponse)
	c.JSON(http.StatusOK, response)

}

func (h *taskController) DeleteTask(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(int)

	var taskDeletedInput input.TaskID
	err := c.ShouldBindUri(&taskDeletedInput)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	taskWillDeleted, err := h.taskService.GetTaskDetail(taskDeletedInput.ID)

	if taskWillDeleted.ID == 0 {
		response := helper.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusNotFound, response)
		return
	}

	if taskWillDeleted.UserID != currentUser {
		response := helper.APIResponse("failed", gin.H{
			"errors": "task not found!",
		})
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	_, err = h.taskService.DeleteTask(taskDeletedInput.ID)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	message := "Task has been successfully deleted"

	response := helper.APIResponse("ok", message)
	c.JSON(http.StatusOK, response)
}
