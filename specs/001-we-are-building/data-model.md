# Data Model: Portfolio Management System

## Entities

### UserProfile
Represents a user's complete professional profile containing multiple sections. Linked to external identity provider via OAuth 2.0 subject identifier.

**Fields**:
- id (UUID) - Primary key
- user_id (string) - Unique OAuth 2.0 subject identifier from external identity provider
- provider_name (string) - Name of the identity provider (e.g., "google", "github", "linkedin")
- created_at (timestamp) - When the profile was created
- updated_at (timestamp) - When the profile was last updated
- is_public (boolean) - Whether the profile is publicly accessible
- view_count (integer) - Number of times the profile has been viewed

**Relationships**:
- Has many ProfileSections
- Has many MediaAssets
- Has many ExternalLinks
- Has one ShareConfiguration

### ProfileSection
Represents an individual section of a profile (summary, education, experience, etc.).

**Fields**:
- id (UUID) - Primary key
- profile_id (UUID) - Foreign key to UserProfile
- type (string) - Type of section (summary, education, experience, skills, etc.)
- title (string) - Title of the section (e.g., "Work Experience", "Education")
- content (text) - Content of the section in structured format
- order (integer) - Display order of sections
- created_at (timestamp) - When the section was created
- updated_at (timestamp) - When the section was last updated

**Relationships**:
- Belongs to UserProfile

### MediaAsset
Represents uploaded media files associated with a profile.

**Fields**:
- id (UUID) - Primary key
- profile_id (UUID) - Foreign key to UserProfile
- filename (string) - Original filename
- storage_key (string) - Key used to retrieve file from NATS object storage
- content_type (string) - MIME type of the file
- size (integer) - Size of the file in bytes
- uploaded_at (timestamp) - When the file was uploaded

**Relationships**:
- Belongs to UserProfile

### ExternalLink
Represents links to external projects associated with a profile.

**Fields**:
- id (UUID) - Primary key
- profile_id (UUID) - Foreign key to UserProfile
- url (string) - URL to external project
- title (string) - Title/description of the link
- added_at (timestamp) - When the link was added

**Relationships**:
- Belongs to UserProfile

### ShareConfiguration
Represents the sharing settings and permissions for a profile.

**Fields**:
- id (UUID) - Primary key
- profile_id (UUID) - Foreign key to UserProfile (one-to-one)
- is_public (boolean) - Whether the profile is publicly accessible
- require_auth (boolean) - Whether authentication is required to view
- allowed_users (JSON) - List of specific users who can view (if not public)
- expiration_date (timestamp) - When the share link expires (optional)
- created_at (timestamp) - When the configuration was created
- updated_at (timestamp) - When the configuration was last updated

**Relationships**:
- Belongs to UserProfile

### ResumeTemplate
Represents the LaTeX template used for generating PDF resumes.

**Fields**:
- id (UUID) - Primary key
- name (string) - Name of the template
- description (text) - Description of the template
- content (text) - LaTeX template content
- is_default (boolean) - Whether this is the default template
- created_at (timestamp) - When the template was created
- updated_at (timestamp) - When the template was last updated

### GeneratedResume
Represents generated PDF resumes.

**Fields**:
- id (UUID) - Primary key
- profile_id (UUID) - Foreign key to UserProfile
- template_id (UUID) - Foreign key to ResumeTemplate
- storage_key (string) - Key used to retrieve PDF from NATS object storage
- generated_at (timestamp) - When the resume was generated
- file_size (integer) - Size of the PDF in bytes

**Relationships**:
- Belongs to UserProfile
- Belongs to ResumeTemplate

## Validation Rules

1. **UserProfile**:
   - user_id must be unique
   - is_public defaults to false

2. **ProfileSection**:
   - type must be one of: summary, education, experience, skills, projects, certifications, publications, volunteer, interests
   - content must be valid JSON for structured data or plain text for simple content

3. **MediaAsset**:
   - content_type must be a valid MIME type
   - size must be positive
   - filename must not be empty

4. **ExternalLink**:
   - url must be a valid URL format
   - title must not be empty

5. **ShareConfiguration**:
   - If is_public is true, allowed_users should be ignored
   - Expiration_date, if set, must be in the future

6. **ResumeTemplate**:
   - content must contain valid LaTeX syntax
   - Only one template can be marked as default

## Indexes

1. UserProfile.user_id (unique)
2. ProfileSection.profile_id + ProfileSection.order (for efficient section retrieval in order)
3. MediaAsset.profile_id (for retrieving all media for a profile)
4. ExternalLink.profile_id (for retrieving all links for a profile)
5. ShareConfiguration.profile_id (unique)
6. GeneratedResume.profile_id (for retrieving resume history)