package models

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// ResumeTemplate holds the schema definition for the ResumeTemplate entity.
type ResumeTemplate struct {
	ent.Schema
}

// Fields of the ResumeTemplate.
func (ResumeTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("Name of the template"),
		field.String("description").
			Optional().
			Comment("Description of the template"),
		field.Text("content").
			NotEmpty().
			Comment("LaTeX template content"),
		field.Bool("is_default").
			Default(false).
			Comment("Whether this is the default template"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the ResumeTemplate.
func (ResumeTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("generated_resumes", GeneratedResume.Type),
	}
}