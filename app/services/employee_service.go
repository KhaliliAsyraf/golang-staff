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

	gender, err := s.GetGender(name)
	if err != nil {
		facades.Log().Error("Failed to get gender, defaulting to unknown:", err)
		gender = "unknown"
	}

	employee := models.Employee{
		Name:          name,
		Email:         email,
		Password:      hashedPassword,
		Type:          "staff",
		IdDepartments: int(departmentId),
		Gender:        gender,
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

func (s *EmployeeService) GetGender(name string) (string, error) {
	genderUrl := facades.Config().Get("http.gender.url")
	res, err := facades.Http().
		WithQueryParameter("name", name).
		Get(genderUrl.(string))
	if err != nil {
		return "", errors.New("something went wrong while fetching employees")
	}

	body, err := res.Json()
	if err != nil {
		return "", errors.New("failed to parse gender response")
	}

	facades.Log().Info("Gender response:", body)

	gender, ok := body["gender"].(string)
	if !ok || gender == "" {
		gender = "male"
	}

	return gender, nil
}
