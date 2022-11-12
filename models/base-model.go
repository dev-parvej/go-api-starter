package models

import "time"

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (model *BaseModel) IsEmpty() bool {
	return model.ID == 0
}
