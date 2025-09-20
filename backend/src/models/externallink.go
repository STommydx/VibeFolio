package models

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// ExternalLink holds the schema definition for the ExternalLink entity.
type ExternalLink struct {
	ent.Schema
}

// Fields of the ExternalLink.
func (ExternalLink) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").
			NotEmpty().
			Comment("URL to external project"),
		field.String("title").
			NotEmpty().
			Comment("Title/description of the link"),
		field.Time("added_at").
			Default(time.Now).
			Comment("When the link was added"),
	}
}

// Edges of the ExternalLink.
func (ExternalLink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_profile", UserProfile.Type).
			Ref("external_links").
			Unique(),
	}
}