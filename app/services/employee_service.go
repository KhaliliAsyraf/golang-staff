package services

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"

	"errors"
)

type EmployeeService struct{}

// Example method
func (s *EmployeeService) StoreEmployee(data map[string]any) (map[string]any, error) {
	email, _ := data["email"].(string)
	name, _ := data["name"].(string)

	hashedPassword, err := facades.Hash().Make("password")
	if err != nil {
		return nil, errors.New("something went wrong while saving employee")
	}

	tx, err := facades.Orm().Query().BeginTransaction()
	if err != nil {
		return nil, errors.New("something went wrong while saving employee")
	}

	employee := models.Employee{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Type:     "staff",
	}
	err = facades.Orm().Query().Create(&employee)

	if err != nil {
		tx.Rollback()
		return nil, errors.New("something went wrong while saving employee")
	}

	tx.Commit()

	return map[string]any{
		"employee": employee,
	}, nil
}
