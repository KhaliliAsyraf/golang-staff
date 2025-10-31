package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251031031154CreateDepartmentsTable struct{}

// Signature The unique signature for the migration.
func (r *M20251031031154CreateDepartmentsTable) Signature() string {
	return "20251031031154_create_departments_table"
}

// Up Run the migrations.
func (r *M20251031031154CreateDepartmentsTable) Up() error {
	if !facades.Schema().HasTable("departments") {
		return facades.Schema().Create("departments", func(table schema.Blueprint) {
			table.ID()
			table.String("name")
			table.TimestampsTz()
			table.SoftDeletes()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20251031031154CreateDepartmentsTable) Down() error {
	return facades.Schema().DropIfExists("departments")
}
