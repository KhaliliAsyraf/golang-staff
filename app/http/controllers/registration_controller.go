package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type RegistrationController struct {
	// Dependent services
}

func NewRegistrationController() *RegistrationController {
	return &RegistrationController{
		// Inject services
	}
}

func (r *RegistrationController) Index(ctx http.Context) http.Response {
	return nil
}

func (r *RegistrationController) Store(ctx http.Context) http.Response {
	// return nil
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Brahh",
	})
}
