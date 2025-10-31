package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type AuthController struct {
	// Dependent services
}

func NewAuthController() *AuthController {
	return &AuthController{
		// Inject services
	}
}

func (r *AuthController) Index(ctx http.Context) http.Response {
	return nil
}

func (r *AuthController) Login(ctx http.Context) http.Response {
	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")

	hashedPassword, err := facades.Hash().Make(password)

	if !facades.Hash().Check(password, hashedPassword) {
		// The passwords match...
		return ctx.Response().Status(500).Json(http.Json{
			"Hello": "Goravel",
		})
	}

	var user models.Employee
	err = facades.
		Orm().
		Query().
		Where("email", email).
		First(&user)
	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"error": "Employee is invalid.",
		})
	}

	token, err := facades.Auth(ctx).Guard("employee").Login(&user)
	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"error": "Failed to login.",
		})
	}
	return ctx.Response().Success().Json(http.Json{
		"token": token,
	})
}
