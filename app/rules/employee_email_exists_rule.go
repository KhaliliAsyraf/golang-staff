package rules

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

type EmployeeEmailExistsRule struct {
}

// Signature The name of the rule.
func (receiver *EmployeeEmailExistsRule) Signature() string {
	return "employee_email_exists"
}

// Passes Determine if the validation rule passes.
func (receiver *EmployeeEmailExistsRule) Passes(data validation.Data, val any, options ...any) bool {
	exists, err := facades.Orm().Query().Model(&models.Employee{}).Where("email", val.(string)).Exists()
	if err != nil {
	}
	return exists
}

// Message Get the validation error message.
func (receiver *EmployeeEmailExistsRule) Message() string {
	return "The :attribute is invalid."
}
