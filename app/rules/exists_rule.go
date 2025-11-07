package rules

import (
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

type ExistsRule struct {
	Table  string
	Column string
}

// Signature The name of the rule.
func (receiver *ExistsRule) Signature() string {
	return "exists"
}

// Passes Determine if the validation rule passes.
func (receiver *ExistsRule) Passes(data validation.Data, val any, options ...any) bool {
	if len(options) < 2 {
		return false
	}
	receiver.Table = options[0].(string)
	receiver.Column = options[1].(string)

	exists, err := facades.Orm().Query().
		Table(receiver.Table).
		Where(receiver.Column, val).
		Exists()
	if err != nil {
	}
	return exists
}

// Message Get the validation error message.
func (receiver *ExistsRule) Message() string {
	return "The :attribute is invalid."
}
