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

/**
 * Index handles the request to get all departments.
 *
 * @param ctx http.Context
 * @return http.Response
 */
func (r *DepartmentController) Index(ctx http.Context) http.Response {
	// return cached
	departments, err := facades.Cache().RememberForever("departments", func() (any, error) {
		var departments []models.Department
		facades.Orm().Query().Get(&departments)
		return departments, nil
	})

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Something wrong with fetching departments",
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"data": departments,
	})
}
