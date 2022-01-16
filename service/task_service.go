package service

import (
	"kanban-golang/model/entity"
	"kanban-golang/model/input"
	"kanban-golang/repository"
)

type TaskService interface {
	CreateTask(input input.TaskCreateInput) (entity.Task, error)
	GetTasks(ID int) ([]entity.Task, error)
	UpdateTask(ID int, editTask input.TaskEditInput) (entity.Task, error)
	UpdateStatusTask(ID int, statusTask input.TaskUpdateStatus) (entity.Task, error)
	DeleteTask(ID int) (bool, error)
}

type taskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) *taskService {
	return &taskService{taskRepository}
}

func (s *taskService) CreateTask(input input.TaskCreateInput) (entity.Task, error) {
	newTask := entity.Task{
		Description: input.Description,
		Title:       input.Title,
		CategoryID:  input.CategoryID,
	}

	createdTask, err := s.taskRepository.Save(newTask)

	if err != nil {
		return createdTask, err
	}

	return createdTask, nil
}

func (s *taskService) GetTasks(ID int) ([]entity.Task, error) {
	tasks, err := s.taskRepository.Get(ID)

	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (s *taskService) UpdateTask(ID int, editTask input.TaskEditInput) (entity.Task, error) {
	// mapping struct
	updateTask := entity.Task{
		Title:       editTask.Title,
		Description: editTask.Description,
	}

	tasks, err := s.taskRepository.Update(ID, updateTask)

	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (s *taskService) UpdateStatusTask(ID int, statusTask input.TaskUpdateStatus) (entity.Task, error) {
	// mapping task
	newStatusTask := statusTask.Status

	taskUpdatedStatus, err := s.taskRepository.SwitchStatus(ID, newStatusTask)

	if err != nil {
		return taskUpdatedStatus, err
	}

	return taskUpdatedStatus, nil
}

func (s *taskService) DeleteTask(ID int) (bool, error) {
	deletedTask, err := s.taskRepository.Delete(ID)

	if err != nil {
		return deletedTask, err
	}

	return deletedTask, nil
}
