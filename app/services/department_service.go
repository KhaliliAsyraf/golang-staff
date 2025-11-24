package services

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type DepartmentService struct{}

func NewDepartmentService() *DepartmentService {
	return &DepartmentService{}
}

func (s *DepartmentService) Create(name string) (*models.Department, error) {
	dept := &models.Department{Name: name}
	if err := facades.Orm().Query().Create(dept); err != nil {
		return nil, err
	}
	return dept, nil
}

func (s *DepartmentService) Find(id int64) (*models.Department, error) {
	var dept models.Department
	if err := facades.Orm().Query().Where("id", id).First(&dept); err != nil {
		return nil, err
	}
	return &dept, nil
}

func (s *DepartmentService) All() ([]models.Department, error) {
	var depts []models.Department
	if err := facades.Orm().Query().Find(&depts); err != nil {
		return nil, err
	}
	return depts, nil
}
