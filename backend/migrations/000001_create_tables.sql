-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE user_profiles (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL UNIQUE,
    provider_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_public BOOLEAN NOT NULL DEFAULT FALSE,
    view_count INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE profile_sections (
    id TEXT PRIMARY KEY,
    profile_id TEXT NOT NULL,
    type TEXT NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    "order" INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (profile_id) REFERENCES user_profiles(id)
);

CREATE TABLE media_assets (
    id TEXT PRIMARY KEY,
    profile_id TEXT NOT NULL,
    filename TEXT NOT NULL,
    storage_key TEXT NOT NULL,
    content_type TEXT NOT NULL,
    size INTEGER NOT NULL,
    uploaded_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (profile_id) REFERENCES user_profiles(id)
);

CREATE TABLE external_links (
    id TEXT PRIMARY KEY,
    profile_id TEXT NOT NULL,
    url TEXT NOT NULL,
    title TEXT NOT NULL,
    added_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (profile_id) REFERENCES user_profiles(id)
);

CREATE TABLE share_configurations (
    id TEXT PRIMARY KEY,
    profile_id TEXT NOT NULL UNIQUE,
    is_public BOOLEAN NOT NULL DEFAULT FALSE,
    require_auth BOOLEAN NOT NULL DEFAULT FALSE,
    allowed_users TEXT,
    expiration_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (profile_id) REFERENCES user_profiles(id)
);

CREATE TABLE resume_templates (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    content TEXT NOT NULL,
    is_default BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE generated_resumes (
    id TEXT PRIMARY KEY,
    profile_id TEXT NOT NULL,
    template_id TEXT NOT NULL,
    storage_key TEXT NOT NULL,
    generated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    file_size INTEGER NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES user_profiles(id),
    FOREIGN KEY (template_id) REFERENCES resume_templates(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE generated_resumes;
DROP TABLE resume_templates;
DROP TABLE share_configurations;
DROP TABLE external_links;
DROP TABLE media_assets;
DROP TABLE profile_sections;
DROP TABLE user_profiles;