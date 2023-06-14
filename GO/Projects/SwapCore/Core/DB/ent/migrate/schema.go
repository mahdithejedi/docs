// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// NetworksColumns holds the columns for the "networks" table.
	NetworksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "network_id", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 120},
		{Name: "symbol", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// NetworksTable holds the schema information for the "networks" table.
	NetworksTable = &schema.Table{
		Name:       "networks",
		Columns:    NetworksColumns,
		PrimaryKey: []*schema.Column{NetworksColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		NetworksTable,
	}
)

func init() {
}
