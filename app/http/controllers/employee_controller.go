package controllers

import (
	"goravel/app/http/requests/employee"
	"goravel/app/services"

	"github.com/goravel/framework/contracts/http"
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

/**
 * Store handles the request to create a new employee.
 *
 * @param ctx http.Context
 * @return http.Response
 */
func (r *EmployeeController) Store(ctx http.Context) http.Response {
	var request employee.StoreEmployeeRequest
	errors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Something wrong with saving employee",
		})
	}

	if errors != nil {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"error": errors.All(),
		})
	}

	// employee := r.employeeService.DoSomething()
	employeeService := services.EmployeeService{}
	var data, storeErr = employeeService.StoreEmployee(ctx.Request().All())
	if storeErr != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": storeErr.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"data": data["employee"],
	})
}
