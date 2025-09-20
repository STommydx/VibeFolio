package components

import (
	"fmt"
	
	"github.com/a-h/templ"
)

// ProfileCard renders a profile card component
func ProfileCard(profileID, name, title, summary string, isPublic bool) templ.Component {
	return templ.ComponentFunc(func(c templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual templ component
		c.Response().Write([]byte("<div class=\"profile-card\">"))
		c.Response().Write([]byte("<h3>" + name + "</h3>"))
		c.Response().Write([]byte("<p>" + title + "</p>"))
		c.Response().Write([]byte("<p>" + summary + "</p>"))
		c.Response().Write([]byte("<p>Public: " + fmt.Sprintf("%t", isPublic) + "</p>"))
		c.Response().Write([]byte("</div>"))
		
		return nil
	})
}