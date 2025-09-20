package models

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// ProfileSection holds the schema definition for the ProfileSection entity.
type ProfileSection struct {
	ent.Schema
}

// Fields of the ProfileSection.
func (ProfileSection) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").
			NotEmpty().
			Comment("Type of section (summary, education, experience, skills, etc.)"),
		field.String("title").
			NotEmpty().
			Comment("Title of the section (e.g., \"Work Experience\", \"Education\")"),
		field.String("content").
			NotEmpty().
			Comment("Content of the section in structured format"),
		field.Int("order").
			NonNegative().
			Comment("Display order of sections"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the ProfileSection.
func (ProfileSection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_profile", UserProfile.Type).
			Ref("profile_sections").
			Unique(),
	}
}