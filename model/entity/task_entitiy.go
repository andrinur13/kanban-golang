package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" valid:"required"`
	Description string    `json:"description" valid:"required"`
	Status      bool      `json:"status"`
	UserID      int       `json:"user_id" valid:"required"`
	CategoryID  int       `json:"category_id" valid:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func (Task) TableName() string {
	return "task"
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(t)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
