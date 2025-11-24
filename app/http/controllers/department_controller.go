package controllers

import (
	"goravel/app/services"

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
	deptService := services.DepartmentService{}
	// return cached
	departments, err := facades.Cache().RememberForever("departments", func() (any, error) {
		return deptService.All()
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
