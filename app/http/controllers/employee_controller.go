package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type EmployeeController struct {
	// Dependent services
}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{
		// Inject services
	}
}

func (r *EmployeeController) Index(ctx http.Context) http.Response {
	return nil
}

func (r *EmployeeController) Store(ctx http.Context) http.Response {
	hashedPassword, err := facades.Hash().Make("password")
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Something wrong with saving employee",
		})
	}

	employee := models.Employee{
		Name:     ctx.Request().Input("name"),
		Email:    ctx.Request().Input("name"),
		Password: hashedPassword,
		Type:     "staff",
	}
	facades.Orm().Query().Create(&employee)

	return ctx.Response().Success().Json(http.Json{
		"message": "Employee created successfully!",
	})
}
