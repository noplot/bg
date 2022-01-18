package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// EntTikTokUser holds the schema definition for the EntTikTokUser entity.
type EntTikTokUser struct {
	ent.Schema
}

// Fields of the EntTikTokUser.
func (EntTikTokUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default(""),
	}
}

// Edges of the EntTikTokUser.
func (EntTikTokUser) Edges() []ent.Edge {
	return nil
}
