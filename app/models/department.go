package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Department struct {
	orm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	// Employees []Employee `gorm:"foreignKey:IdDepartments;references:ID"`
}
