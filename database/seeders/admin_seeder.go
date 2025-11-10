package seeders

import (
	"fmt"

	"github.com/goravel/framework/facades"

	"goravel/app/models"
)

type AdminSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *AdminSeeder) Signature() string {
	return "AdminSeeder"
}

// Run executes the seeder logic.
func (s *AdminSeeder) Run() error {
	// return nil
	hashedPassword, err := facades.Hash().Make("password")
	if err != nil {
		// Handle the "catch" part
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user := []models.Employee{
		{
			Name:     "admin",
			Email:    "admin@example.com",
			Password: hashedPassword,
			Type:     "admin",
			Gender:   "male",
		},
	}
	return facades.Orm().Query().Create(&user)
}
