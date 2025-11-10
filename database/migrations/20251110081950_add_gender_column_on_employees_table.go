package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251110081950AddGenderColumnOnEmployeesTable struct{}

// Signature The unique signature for the migration.
func (r *M20251110081950AddGenderColumnOnEmployeesTable) Signature() string {
	return "20251110081950_add_gender_column_on_employees_table"
}

// Up Run the migrations.
func (r *M20251110081950AddGenderColumnOnEmployeesTable) Up() error {
	facades.Schema().Table("employees", func(table schema.Blueprint) {
		table.Enum("gender", []any{"male", "female"})
	})
	return nil
}

// Down Reverse the migrations.
func (r *M20251110081950AddGenderColumnOnEmployeesTable) Down() error {
	facades.Schema().Table("employees", func(table schema.Blueprint) {
		table.DropColumn("gender")
	})
	return nil
}
