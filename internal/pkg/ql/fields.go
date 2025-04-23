package ql

import (
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	"github.com/samber/lo"
)

type (
	// Field describes column of a db table
	Field interface {
		// AsText returns column value as text like "table.column::TEXT AS table_column"
		AsText() string
		// Full returns full column name as "table.column"
		Full() string
		// Short returns short column name, as "column"
		Short() string
		// AS returns full column name with an alias, like "table.column AS table_column"
		AS() string
	}

	tableField struct {
		table  string
		column string
	}
)

func (t tableField) AsText() string {
	return t.Full() + "::TEXT AS " + t.Alias()
}

func (t tableField) Full() string {
	return t.table + "." + t.column
}

func (t tableField) Short() string {
	return t.column
}

func (t tableField) Alias() string {
	return t.table + "_" + t.column
}

func (t tableField) AS() string {
	return t.Full() + " AS " + t.Alias()
}

// Common functions

// ON returns join expression "table_a.a = table_b.b"
// For use like "Join(TableName, table.ON(id, referenced_id))"
func ON(a, b Field) string {
	return a.Full() + " = " + b.Full()
}

// NOT returns "NOT field" expression.
func NOT(field Field) string {
	return "NOT " + field.Full()
}

// NOW returns "NOW()" sql function.
func NOW() string {
	return "NOW()"
}

// CountAll returns 'COUNT(*)' aggregation
func CountAll() string {
	return "COUNT(*)"
}

func DistinctOn(field Field) string {
	return fmt.Sprintf("DISTINCT ON(%s) %s", field.Full(), field.Full())
}

func Eq(a, b Field) string {
	return fmt.Sprintf("%s = %s", a.Full(), b.Full())
}

// Count returns 'COUNT(field)' aggregation
func Count(field string) string {
	return fmt.Sprintf("COUNT(%s)", field)
}

// Distinct returns 'Distinct field' aggregation
func Distinct(field string) string {
	return fmt.Sprintf("DISTINCT %s", field)
}

// NewField creates Field from table and column values.
func NewField(table, column string) Field {
	return tableField{
		table:  table,
		column: column,
	}
}

type Fields []Field

func (fs Fields) Shorts() []string {
	return lo.Map(fs, func(f Field, _ int) string {
		return f.Short()
	})
}

func (fs Fields) ToAssignments(builder *sqlbuilder.UpdateBuilder, values ...any) []string {
	var assignments = make([]string, len(fs))
	for i, field := range fs {
		assignments[i] = builder.Assign(field.Short(), values[i])
	}

	return assignments
}
