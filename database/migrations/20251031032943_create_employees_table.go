package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251031032943CreateEmployeesTable struct{}

// Signature The unique signature for the migration.
func (r *M20251031032943CreateEmployeesTable) Signature() string {
	return "20251031032943_create_employees_table"
}

// Up Run the migrations.
func (r *M20251031032943CreateEmployeesTable) Up() error {
	if !facades.Schema().HasTable("employees") {
		return facades.Schema().Create("employees", func(table schema.Blueprint) {
			table.ID("id")
			table.String("name")
			table.String("email")
			table.String("password")
			table.Enum("type", []any{"admin", "staff"})
			table.TimestampsTz()
			table.SoftDeletes()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20251031032943CreateEmployeesTable) Down() error {
	return facades.Schema().DropIfExists("employees")
}
