package repository

import (
	"kanban-golang/model/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Save(task entity.Task) (entity.Task, error)
	Get(IDUser int) ([]entity.Task, error)
	GetDetail(ID int) (entity.Task, error)
	Update(ID int, taskEdit entity.Task) (entity.Task, error)
	SwitchStatus(ID int, status bool) (entity.Task, error)
	Delete(ID int) (bool, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Save(task entity.Task) (entity.Task, error) {
	err := r.db.Create(&task).Error

	if err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) Get(IDUser int) ([]entity.Task, error) {
	allTask := []entity.Task{}

	err := r.db.Where("user_id = ?", IDUser).Find(&allTask).Error

	if err != nil {
		return []entity.Task{}, err
	}

	return allTask, nil

}

func (r *taskRepository) GetDetail(ID int) (entity.Task, error) {
	task := entity.Task{}

	err := r.db.Where("id = ?", ID).Find(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *taskRepository) Update(ID int, taskEdit entity.Task) (entity.Task, error) {
	err := r.db.Where("id = ?", ID).Updates(taskEdit).Error

	if err != nil {
		return entity.Task{}, err
	}

	updatedTask := entity.Task{}
	err = r.db.Where("id = ?", ID).Find(&updatedTask).Error

	if err != nil {
		return entity.Task{}, err
	}

	return updatedTask, nil
}

func (r *taskRepository) SwitchStatus(ID int, status bool) (entity.Task, error) {
	err := r.db.Where("id = ?", ID).Updates(entity.Task{Status: status}).Error

	if err != nil {
		return entity.Task{}, err
	}

	taskSwitched := entity.Task{}
	err = r.db.Where("id = ?", ID).Find(&taskSwitched).Error

	if err != nil {
		return entity.Task{}, err
	}

	return taskSwitched, nil
}

func (r *taskRepository) Delete(ID int) (bool, error) {

	taskDeleted := entity.Task{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(&taskDeleted).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
