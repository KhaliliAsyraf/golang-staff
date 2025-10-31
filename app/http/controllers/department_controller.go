package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type DepartmentController struct {
	// Dependent services
}

func NewDepartmentController() *DepartmentController {
	return &DepartmentController{
		// Inject services
	}
}

func (r *DepartmentController) Index(ctx http.Context) http.Response {
	var departments []models.Department
	facades.Orm().Query().Get(&departments)

	return ctx.Response().Json(http.StatusOK, http.Json{
		"data": departments,
	})
}
