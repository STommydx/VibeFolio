package models

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// GeneratedResume holds the schema definition for the GeneratedResume entity.
type GeneratedResume struct {
	ent.Schema
}

// Fields of the GeneratedResume.
func (GeneratedResume) Fields() []ent.Field {
	return []ent.Field{
		field.String("storage_key").
			NotEmpty().
			Comment("Key used to retrieve PDF from NATS object storage"),
		field.Time("generated_at").
			Default(time.Now).
			Comment("When the resume was generated"),
		field.Int("file_size").
			Positive().
			Comment("Size of the PDF in bytes"),
	}
}

// Edges of the GeneratedResume.
func (GeneratedResume) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_profile", UserProfile.Type).
			Ref("generated_resumes").
			Unique(),
		edge.From("resume_template", ResumeTemplate.Type).
			Ref("generated_resumes").
			Unique(),
	}
}