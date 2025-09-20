# Quickstart Guide: Portfolio Management System

## Overview
This guide provides step-by-step instructions for setting up and using the Portfolio Management System.

## Prerequisites
- Go 1.22 or later
- LaTeX distribution (for PDF generation)
- Git

## Setup
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd VibeFolio
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the application:
   ```bash
   go build -o vibefolio .
   ```

4. Run the application:
   ```bash
   ./vibefolio serve
   ```

## Quickstart Steps

### 1. Create a User Profile
- Navigate to `http://localhost:8080/profiles/new`
- Fill in the profile details
- Set privacy settings (public/private)
- Click "Create Profile"

### 2. Add Profile Sections
- From your profile page, click "Add Section"
- Choose section type (summary, education, experience, etc.)
- Fill in the section content
- Save the section

### 3. Upload Media
- From your profile page, click "Upload Media"
- Select files to upload
- Add descriptions for each file
- Click "Upload"

### 4. Add External Links
- From your profile page, click "Add Link"
- Enter the URL and description
- Click "Save Link"

### 5. Generate a Resume
- From your profile page, click "Generate Resume"
- Select a template (if multiple available)
- Click "Generate"
- Wait for the generation to complete (you'll receive a notification)
- Download the generated PDF

### 6. Share Your Profile
- From your profile page, click "Share Settings"
- Configure who can view your profile:
  - Public (anyone with the link)
  - Private with specific users
  - Private (only you)
- Set expiration date if desired
- Save settings
- Copy the share link to share with others

## API Usage Examples

### Create a Profile (CLI)
```bash
vibefolio profiles create --public=true
```

### Add a Section (CLI)
```bash
vibefolio profiles sections add --profile-id=PROFILE_ID --type=summary --title="Professional Summary" --content="Experienced software developer..."
```

### Generate a Resume (CLI)
```bash
vibefolio profiles resume generate --profile-id=PROFILE_ID
```

## Configuration
The application can be configured using a HCL configuration file located at:
- Linux: `$XDG_CONFIG_HOME/vibefolio/config.hcl` or `~/.config/vibefolio/config.hcl`
- macOS: `~/Library/Application Support/vibefolio/config.hcl`
- Windows: `%APPDATA%\vibefolio\config.hcl`

Example configuration:
```hcl
server {
  port = 8080
  host = "localhost"
}

database {
  driver = "sqlite"
  dsn = "./data/vibefolio.db"
}

nats {
  url = "nats://localhost:4222"
  embedded = true
}

logging {
  level = "info"
  format = "json"
}
```

## Troubleshooting
1. **PDF Generation Fails**: Ensure LaTeX is properly installed and accessible in PATH
2. **Media Upload Fails**: Check file size limits and supported formats
3. **Profile Not Accessible**: Verify share settings and authentication status

## Next Steps
- Explore advanced profile customization options
- Create multiple profiles for different purposes
- Set up automated resume generation
- Integrate with external portfolio platforms