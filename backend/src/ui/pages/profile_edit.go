package pages

import (
	"fmt"
	
	"github.com/a-h/templ"
	"github.com/stommydx/VibeFolio/backend/src/ui/components"
)

// ProfileEditPage renders the profile edit page
func ProfileEditPage(profile ProfileDetail) templ.Component {
	return templ.ComponentFunc(func(c templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual templ page
		c.Response().Write([]byte("<html>"))
		c.Response().Write([]byte("<head><title>Edit Profile</title></head>"))
		c.Response().Write([]byte("<body>"))
		c.Response().Write([]byte("<h1>Edit Profile</h1>"))
		
		// Render the profile form
		c.Response().Write([]byte("<form method=\"POST\" action=\"/profiles/" + profile.ID + "\">"))
		c.Response().Write([]byte("<input type=\"hidden\" name=\"_method\" value=\"PUT\">"))
		c.Response().Write([]byte("<div>"))
		c.Response().Write([]byte("<label for=\"name\">Name:</label>"))
		c.Response().Write([]byte("<input type=\"text\" id=\"name\" name=\"name\" value=\"" + profile.Name + "\">"))
		c.Response().Write([]byte("</div>"))
		
		c.Response().Write([]byte("<div>"))
		c.Response().Write([]byte("<label for=\"title\">Title:</label>"))
		c.Response().Write([]byte("<input type=\"text\" id=\"title\" name=\"title\" value=\"" + profile.Title + "\">"))
		c.Response().Write([]byte("</div>"))
		
		c.Response().Write([]byte("<div>"))
		c.Response().Write([]byte("<label for=\"summary\">Summary:</label>"))
		c.Response().Write([]byte("<textarea id=\"summary\" name=\"summary\">" + profile.Summary + "</textarea>"))
		c.Response().Write([]byte("</div>"))
		
		c.Response().Write([]byte("<div>"))
		c.Response().Write([]byte("<label>"))
		c.Response().Write([]byte("<input type=\"checkbox\" id=\"is_public\" name=\"is_public\" " + checked(profile.IsPublic) + ">"))
		c.Response().Write([]byte("<span>Public Profile</span>"))
		c.Response().Write([]byte("</label>"))
		c.Response().Write([]byte("</div>"))
		
		c.Response().Write([]byte("<button type=\"submit\">Save Changes</button>"))
		c.Response().Write([]byte("</form>"))
		
		c.Response().Write([]byte("<a href=\"/profiles/" + profile.ID + "\">Cancel</a>"))
		c.Response().Write([]byte("</body>"))
		c.Response().Write([]byte("</html>"))
		
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