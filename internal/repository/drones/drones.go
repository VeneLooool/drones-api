package drones

import "github.com/VeneLooool/drones-api/internal/pkg/ql"

const Table = "drones"

var (
	ID     = ql.NewField(Table, "id")
	Name   = ql.NewField(Table, "name")
	Status = ql.NewField(Table, "status")
)
