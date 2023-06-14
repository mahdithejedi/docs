package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Network holds the schema definition for the Network entity.
type Network struct {
	ent.Schema
}

// Fields of the Network.
func (Network) Fields() []ent.Field {
	return []ent.Field{
		field.Int("NetworkID").Positive(),
		field.String("name").MaxLen(120).Unique().Validate(UpperCaseValidator),
		field.String("symbol").MaxLen(20).Unique().Validate(UpperCaseValidator),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Network.
func (Network) Edges() []ent.Edge {
	return nil
}
