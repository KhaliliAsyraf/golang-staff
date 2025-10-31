package models

import (
	"github.com/goravel/framework/database/orm"
)

type Employee struct {
	orm.Model
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
	Type     string
}

func (r *Employee) TableName() string {
	return "employees"
}
