// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SamplesColumns holds the columns for the "samples" table.
	SamplesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Default: "sample"},
	}
	// SamplesTable holds the schema information for the "samples" table.
	SamplesTable = &schema.Table{
		Name:       "samples",
		Columns:    SamplesColumns,
		PrimaryKey: []*schema.Column{SamplesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SamplesTable,
	}
)

func init() {
}
