package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
	"goravel/app/http/middleware"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)

	authController := controllers.NewAuthController()
	facades.Route().Post("/login", authController.Login)

	registrationController := controllers.NewRegistrationController()
	facades.Route().Post("/registration", registrationController.Store)

	// --- Employee ---
	employeeController := controllers.NewEmployeeController()
	facades.Route().
		Middleware(middleware.Auth()).
		Prefix("employee").
		Group(func(router route.Router) {
			router.Post("/", employeeController.Store)
		})

	// --- Department ---
	departmentController := controllers.NewDepartmentController()
	facades.Route().
		Prefix("department").
		Middleware(middleware.Auth()).
		Group(func(router route.Router) {
			router.Get("/", departmentController.Index)
		})
}
