package employee

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type StoreEmployeeRequest struct {
	Name         string `form:"name" json:"name"`
	Email        string `form:"email" json:"email"`
	DepartmentId int    `form:"department_id" json:"department_id"`
}

func (r *StoreEmployeeRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreEmployeeRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StoreEmployeeRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":          "required|max_len:255",
		"email":         "required|email",
		"department_id": "required|exists:departments,id",
	}
}

func (r *StoreEmployeeRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StoreEmployeeRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StoreEmployeeRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
