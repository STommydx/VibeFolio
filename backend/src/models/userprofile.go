package models

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// UserProfile holds the schema definition for the UserProfile entity.
type UserProfile struct {
	ent.Schema
}

// Fields of the UserProfile.
func (UserProfile) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			NotEmpty().
			Unique().
			Comment("Unique OAuth 2.0 subject identifier from external identity provider"),
		field.String("provider_name").
			NotEmpty().
			Comment("Name of the identity provider (e.g., \"google\", \"github\", \"linkedin\")"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Bool("is_public").
			Default(false).
			Comment("Whether the profile is publicly accessible"),
		field.Int("view_count").
			NonNegative().
			Default(0).
			Comment("Number of times the profile has been viewed"),
	}
}

// Edges of the UserProfile.
func (UserProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile_sections", ProfileSection.Type),
		edge.To("media_assets", MediaAsset.Type),
		edge.To("external_links", ExternalLink.Type),
		edge.To("share_configuration", ShareConfiguration.Type).
			Unique(),
	}
}