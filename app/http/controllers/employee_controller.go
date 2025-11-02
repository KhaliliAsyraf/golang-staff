package controllers

import (
	"goravel/app/http/requests/employee"
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

	hashedPassword, err := facades.Hash().Make("password")
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Something wrong with saving employee",
		})
	}

	tx, err := facades.Orm().Query().BeginTransaction()
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Something wrong with saving employee",
		})
	}

	employee := models.Employee{
		Name:     ctx.Request().Input("name"),
		Email:    ctx.Request().Input("email"),
		Password: hashedPassword,
		Type:     "staff",
	}
	err = facades.Orm().Query().Create(&employee)

	if err != nil {
		tx.Rollback()
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Something wrong with saving employee",
		})
	}

	tx.Commit()

	return ctx.Response().Success().Json(http.Json{
		"message": "Employee created successfully!",
	})
}
