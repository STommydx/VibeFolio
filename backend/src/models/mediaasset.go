package models

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// MediaAsset holds the schema definition for the MediaAsset entity.
type MediaAsset struct {
	ent.Schema
}

// Fields of the MediaAsset.
func (MediaAsset) Fields() []ent.Field {
	return []ent.Field{
		field.String("filename").
			NotEmpty().
			Comment("Original filename"),
		field.String("storage_key").
			NotEmpty().
			Comment("Key used to retrieve file from NATS object storage"),
		field.String("content_type").
			NotEmpty().
			Comment("MIME type of the file"),
		field.Int("size").
			Positive().
			Comment("Size of the file in bytes"),
		field.Time("uploaded_at").
			Default(time.Now).
			Comment("When the file was uploaded"),
	}
}

// Edges of the MediaAsset.
func (MediaAsset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_profile", UserProfile.Type).
			Ref("media_assets").
			Unique(),
	}
}