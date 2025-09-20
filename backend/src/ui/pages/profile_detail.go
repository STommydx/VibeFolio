package pages

import (
	"fmt"
	
	"github.com/a-h/templ"
	"github.com/stommydx/VibeFolio/backend/src/ui/components"
)

// ProfileDetailPage renders the profile detail page
func ProfileDetailPage(profile ProfileDetail, sections []ProfileSection) templ.Component {
	return templ.ComponentFunc(func(c templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual templ page
		c.Response().Write([]byte("<html>"))
		c.Response().Write([]byte("<head><title>" + profile.Name + "</title></head>"))
		c.Response().Write([]byte("<body>"))
		c.Response().Write([]byte("<h1>" + profile.Name + "</h1>"))
		c.Response().Write([]byte("<p>Title: " + profile.Title + "</p>"))
		c.Response().Write([]byte("<p>Summary: " + profile.Summary + "</p>"))
		c.Response().Write([]byte("<p>Public: " + fmt.Sprintf("%t", profile.IsPublic) + "</p>"))
		
		c.Response().Write([]byte("<h2>Sections</h2>"))
		for _, section := range sections {
			c.Response().Write([]byte("<div>"))
			c.Response().Write([]byte("<h3>" + section.Title + "</h3>"))
			c.Response().Write([]byte("<p>Type: " + section.Type + "</p>"))
			c.Response().Write([]byte("<div>" + section.Content + "</div>"))
			c.Response().Write([]byte("</div>"))
		}
		
		c.Response().Write([]byte("<a href=\"/profiles/" + profile.ID + "/edit\">Edit Profile</a>"))
		c.Response().Write([]byte("<a href=\"/profiles\">Back to List</a>"))
		c.Response().Write([]byte("</body>"))
		c.Response().Write([]byte("</html>"))
		
		return nil
	})
}

// ProfileDetail represents detailed profile information
type ProfileDetail struct {
	ID       string
	Name     string
	Title    string
	Summary  string
	IsPublic bool
}

// ProfileSection represents a profile section
type ProfileSection struct {
	ID      string
	Type    string
	Title   string
	Content string
}