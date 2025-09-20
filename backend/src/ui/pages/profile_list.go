package pages

import (
	"fmt"
	
	"github.com/a-h/templ"
	"github.com/stommydx/VibeFolio/backend/src/ui/components"
)

// ProfileListPage renders the profile list page
func ProfileListPage(profiles []Profile) templ.Component {
	return templ.ComponentFunc(func(c templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual templ page
		c.Response().Write([]byte("<html>"))
		c.Response().Write([]byte("<head><title>Profile List</title></head>"))
		c.Response().Write([]byte("<body>"))
		c.Response().Write([]byte("<h1>Profiles</h1>"))
		c.Response().Write([]byte("<ul>"))
		
		for _, profile := range profiles {
			c.Response().Write([]byte("<li>"))
			c.Response().Write([]byte("<a href=\"/profiles/" + profile.ID + "\">" + profile.Name + "</a>"))
			c.Response().Write([]byte("</li>"))
		}
		
		c.Response().Write([]byte("</ul>"))
		c.Response().Write([]byte("<a href=\"/profiles/new\">Create New Profile</a>"))
		c.Response().Write([]byte("</body>"))
		c.Response().Write([]byte("</html>"))
		
		return nil
	})
}

// Profile represents a user profile
type Profile struct {
	ID   string
	Name string
}