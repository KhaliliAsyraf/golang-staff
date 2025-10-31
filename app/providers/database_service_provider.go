package providers

import (
	"github.com/goravel/framework/contracts/database/seeder"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"

	"goravel/database"
	"goravel/database/seeders"
)

type DatabaseServiceProvider struct {
}

func (receiver *DatabaseServiceProvider) Register(app foundation.Application) {
}

func (receiver *DatabaseServiceProvider) Boot(app foundation.Application) {
	kernel := database.Kernel{}
	facades.Schema().Register(kernel.Migrations())
	// facades.Seeder().Register(kernel.Seeders())

	facades.Seeder().Register([]seeder.Seeder{
		// &seeders.DatabaseSeeder{},
		&seeders.AdminSeeder{},
		&seeders.DepartmentSeeder{},
	})
}
