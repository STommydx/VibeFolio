package components

import (
	"github.com/a-h/templ"
)

// ProfileSection renders a profile section component
func ProfileSection(sectionID, sectionType, title, content string) templ.Component {
	return templ.ComponentFunc(func(ctx templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual templ component
		ctx.ResponseWriter().Write([]byte("<div class=\"profile-section\">"))
		ctx.ResponseWriter().Write([]byte("<h4>" + title + "</h4>"))
		ctx.ResponseWriter().Write([]byte("<p>Type: " + sectionType + "</p>"))
		ctx.ResponseWriter().Write([]byte("<div>" + content + "</div>"))
		ctx.ResponseWriter().Write([]byte("</div>"))
		
		return nil
	})
}