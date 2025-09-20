package components

import (
	"fmt"
	
	"github.com/a-h/templ"
)

// ProfileForm renders a profile form component
func ProfileForm(profileID, name, title, summary string, isPublic bool) templ.Component {
	return templ.ComponentFunc(func(c templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual templ component
		c.Response().Write([]byte("<form class=\"profile-form\">"))
		c.Response().Write([]byte("<input type=\"text\" name=\"name\" value=\"" + name + "\" placeholder=\"Name\">"))
		c.Response().Write([]byte("<input type=\"text\" name=\"title\" value=\"" + title + "\" placeholder=\"Title\">"))
		c.Response().Write([]byte("<textarea name=\"summary\" placeholder=\"Summary\">" + summary + "</textarea>"))
		c.Response().Write([]byte("<label><input type=\"checkbox\" name=\"is_public\" " + checked(isPublic) + "> Public</label>"))
		c.Response().Write([]byte("<button type=\"submit\">Save Profile</button>"))
		c.Response().Write([]byte("</form>"))
		
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