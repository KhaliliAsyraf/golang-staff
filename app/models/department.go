package models

import (
	"github.com/goravel/framework/database/orm"
)

type Department struct {
	orm.Model
	Name string
}
