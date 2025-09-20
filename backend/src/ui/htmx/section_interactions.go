package htmx

import (
	"fmt"
	
	"github.com/a-h/templ"
)

// AddSectionForm renders a form for adding a new section
func AddSectionForm(profileID string) templ.Component {
	return templ.ComponentFunc(func(c templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual HTMX component
		c.Response().Write([]byte("<form hx-post=\"/profiles/" + profileID + "/sections\" hx-target=\"#sections-list\" hx-swap=\"beforeend\">"))
		c.Response().Write([]byte("<select name=\"type\">"))
		c.Response().Write([]byte("<option value=\"summary\">Summary</option>"))
		c.Response().Write([]byte("<option value=\"education\">Education</option>"))
		c.Response().Write([]byte("<option value=\"experience\">Experience</option>"))
		c.Response().Write([]byte("<option value=\"skills\">Skills</option>"))
		c.Response().Write([]byte("</select>"))
		c.Response().Write([]byte("<input type=\"text\" name=\"title\" placeholder=\"Section Title\">"))
		c.Response().Write([]byte("<textarea name=\"content\" placeholder=\"Section Content\"></textarea>"))
		c.Response().Write([]byte("<button type=\"submit\">Add Section</button>"))
		c.Response().Write([]byte("</form>"))
		
		return nil
	})
}

// SectionItem renders a section item with HTMX interactions
func SectionItem(sectionID, profileID, title, content string) templ.Component {
	return templ.ComponentFunc(func(c templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual HTMX component
		c.Response().Write([]byte("<div id=\"section-" + sectionID + "\" class=\"section-item\">"))
		c.Response().Write([]byte("<h3>" + title + "</h3>"))
		c.Response().Write([]byte("<div>" + content + "</div>"))
		c.Response().Write([]byte("<button hx-delete=\"/profiles/" + profileID + "/sections/" + sectionID + "\" hx-target=\"#section-" + sectionID + "\" hx-swap=\"outerHTML\">Delete</button>"))
		c.Response().Write([]byte("</div>"))
		
		return nil
	})
}

// DeleteSectionButton renders a delete button for a section
func DeleteSectionButton(sectionID, profileID string) templ.Component {
	return templ.ComponentFunc(func(c templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual HTMX component
		c.Response().Write([]byte("<button hx-delete=\"/profiles/" + profileID + "/sections/" + sectionID + "\" hx-target=\"#section-" + sectionID + "\" hx-swap=\"outerHTML\">Delete</button>"))
		
		return nil
	})
}