package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"

	customMiddleware "goravel/app/http/middleware"

	importMiddleware "github.com/goravel/framework/http/middleware"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)

	authController := controllers.NewAuthController()
	facades.Route().
		Middleware(importMiddleware.Throttle("login")). // rate limiting middleware
		Post("/login", authController.Login)

	// --- Employee ---
	employeeController := controllers.NewEmployeeController()
	facades.Route().
		Middleware(customMiddleware.Auth()).
		Prefix("employee").
		Group(func(router route.Router) {
			router.Get("/", employeeController.Index)
			router.Post("/", employeeController.Store)
		})

	// --- Department ---
	departmentController := controllers.NewDepartmentController()
	facades.Route().
		Prefix("department").
		Middleware(customMiddleware.Auth()).
		Group(func(router route.Router) {
			router.Get("/", departmentController.Index)
		})
}
