package models

import (
	"github.com/goravel/framework/database/orm"
)

type Department struct {
	orm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
	// Employees []Employee `gorm:"foreignKey:IdDepartments;references:ID"`
}
