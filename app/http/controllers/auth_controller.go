package controllers

import (
	"goravel/app/http/requests/auth"
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
	var request auth.LoginRequest
	errors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Something wrong with logging in",
		})
	}

	if errors != nil {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"error": errors.All(),
		})
	}

	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")

	hashedPassword, err := facades.Hash().Make(password)

	if !facades.Hash().Check(password, hashedPassword) {
		// The passwords match...
		return ctx.Response().Status(401).Json("Unauthorized")
	}

	var user models.Employee
	err = facades.
		Orm().
		Query().
		Where("email", email).
		First(&user)
	if err != nil {
		return ctx.Response().Status(400).Json(http.Json{
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
