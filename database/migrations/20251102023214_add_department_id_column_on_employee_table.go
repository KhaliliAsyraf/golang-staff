package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251102023214AddDepartmentIdColumnOnEmployeeTable struct{}

// Signature The unique signature for the migration.
func (r *M20251102023214AddDepartmentIdColumnOnEmployeeTable) Signature() string {
	return "20251102023214_add_department_id_column_on_employee_table"
}

// Up Run the migrations.
func (r *M20251102023214AddDepartmentIdColumnOnEmployeeTable) Up() error {
	facades.Schema().Table("employees", func(table schema.Blueprint) {
		table.Integer("id_departments").Nullable().After("email")
	})
	return nil
}

// Down Reverse the migrations.
func (r *M20251102023214AddDepartmentIdColumnOnEmployeeTable) Down() error {
	facades.Schema().Table("employees", func(table schema.Blueprint) {
		table.DropColumn("id_departments")
	})
	return nil
}
