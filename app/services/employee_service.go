package services

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"

	"errors"
)

type EmployeeService struct{}

func (s *EmployeeService) StoreEmployee(data map[string]any) (map[string]any, error) {
	email, _ := data["email"].(string)
	name, _ := data["name"].(string)
	departmentId, _ := data["department_id"].(float64) // JSON numbers are float64

	hashedPassword, err := facades.Hash().Make("password")
	if err != nil {
		return nil, errors.New("something went wrong while saving employee")
	}

	tx, err := facades.Orm().Query().BeginTransaction()
	if err != nil {
		return nil, errors.New("something went wrong while saving employee")
	}

	facades.Log().Info("Dept ID: ", departmentId)

	employee := models.Employee{
		Name:          name,
		Email:         email,
		Password:      hashedPassword,
		Type:          "staff",
		IdDepartments: int(departmentId),
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

func (s *EmployeeService) List(data map[string]any) (map[string]any, error) {
	var employees []models.Employee
	err := facades.Orm().Query().With("Department").Get(&employees)
	if err != nil {
		return nil, errors.New("something went wrong while fetching employees")
	}

	return map[string]any{
		"employees": employees,
	}, nil
}
