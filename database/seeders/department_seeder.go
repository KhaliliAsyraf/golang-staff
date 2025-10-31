package seeders

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type DepartmentSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *DepartmentSeeder) Signature() string {
	return "DepartmentSeeder"
}

// Run executes the seeder logic.
func (s *DepartmentSeeder) Run() error {
	dpmt := []models.Department{
		{
			Name: "HR",
		},
		{
			Name: "IT",
		},
		{
			Name: "Finance",
		},
		{
			Name: "Risk",
		},
	}
	return facades.Orm().Query().Create(&dpmt)
}
