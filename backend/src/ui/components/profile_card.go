package components

import (
	"github.com/a-h/templ"
)

// ProfileCard renders a profile card component
func ProfileCard(profileID, name, title, summary string, isPublic bool) templ.Component {
	return templ.ComponentFunc(func(ctx templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual templ component
		ctx.ResponseWriter().Write([]byte("<div class=\"profile-card\">"))
		ctx.ResponseWriter().Write([]byte("<h3>" + name + "</h3>"))
		ctx.ResponseWriter().Write([]byte("<p>" + title + "</p>"))
		ctx.ResponseWriter().Write([]byte("<p>" + summary + "</p>"))
		ctx.ResponseWriter().Write([]byte("<p>Public: " + checked(isPublic) + "</p>"))
		ctx.ResponseWriter().Write([]byte("</div>"))
		
		return nil
	})
}

// Helper function to generate checked attribute
func checked(value bool) string {
	if value {
		return "checked"
	}
	return ""
}