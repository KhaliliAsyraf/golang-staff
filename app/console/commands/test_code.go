package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
)

type TestCode struct {
}

// Signature The name and signature of the console command.
func (r *TestCode) Signature() string {
	return "app:test"
}

// Description The console command description.
func (r *TestCode) Description() string {
	return "Command description"
}

// Extend The console command extend.
func (r *TestCode) Extend() command.Extend {
	return command.Extend{Category: "app"}
}

// Handle Execute the console command.
func (r *TestCode) Handle(ctx console.Context) error {
	ctx.Info("halu")
	return nil
}
