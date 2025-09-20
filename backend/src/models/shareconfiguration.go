package models

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// ShareConfiguration holds the schema definition for the ShareConfiguration entity.
type ShareConfiguration struct {
	ent.Schema
}

// Fields of the ShareConfiguration.
func (ShareConfiguration) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_public").
			Default(false).
			Comment("Whether the profile is publicly accessible"),
		field.Bool("require_auth").
			Default(false).
			Comment("Whether authentication is required to view"),
		field.JSON("allowed_users", []string{}).
			Optional().
			Comment("List of specific users who can view (if not public)"),
		field.Time("expiration_date").
			Optional().
			Comment("When the share link expires (optional)"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the ShareConfiguration.
func (ShareConfiguration) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_profile", UserProfile.Type).
			Ref("share_configuration").
			Unique(),
	}
}