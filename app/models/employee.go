package models

import (
	"github.com/goravel/framework/database/orm"
)

type Employee struct {
	orm.Model
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string `json:"-"`
	Type     string `gorm:"type:enum('admin','staff');default:'staff'"`
	Gender   string

	IdDepartments int        `gorm:"column:id_departments"`
	Department    Department `gorm:"foreignKey:IdDepartments;references:ID"`
}

func (r *Employee) TableName() string {
	return "employees"
}
