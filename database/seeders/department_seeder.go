package seeders

import (
	"goravel/app/models"
	"log"

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
	err := facades.DB().Statement("TRUNCATE TABLE departments")
	if err != nil {
		log.Println("truncate error:", err)
		return err
	}

	dpmt := []models.Department{
		{
			Name: "Admin",
		},
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

	if err := facades.Orm().Query().Create(&dpmt); err != nil {
		return err
	}

	if err := facades.Cache().Put("departments", dpmt, 0); err != nil {
		return err
	}

	return nil
}
