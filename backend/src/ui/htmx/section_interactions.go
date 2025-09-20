package htmx

import (
	"github.com/a-h/templ"
)

// AddSectionForm renders a form for adding a new section
func AddSectionForm(profileID string) templ.Component {
	return templ.ComponentFunc(func(ctx templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual HTMX component
		ctx.ResponseWriter().Write([]byte("<form hx-post=\"/profiles/" + profileID + "/sections\" hx-target=\"#sections-list\" hx-swap=\"beforeend\">"))
		ctx.ResponseWriter().Write([]byte("<select name=\"type\">"))
		ctx.ResponseWriter().Write([]byte("<option value=\"summary\">Summary</option>"))
		ctx.ResponseWriter().Write([]byte("<option value=\"education\">Education</option>"))
		ctx.ResponseWriter().Write([]byte("<option value=\"experience\">Experience</option>"))
		ctx.ResponseWriter().Write([]byte("<option value=\"skills\">Skills</option>"))
		ctx.ResponseWriter().Write([]byte("</select>"))
		ctx.ResponseWriter().Write([]byte("<input type=\"text\" name=\"title\" placeholder=\"Section Title\">"))
		ctx.ResponseWriter().Write([]byte("<textarea name=\"content\" placeholder=\"Section Content\"></textarea>"))
		ctx.ResponseWriter().Write([]byte("<button type=\"submit\">Add Section</button>"))
		ctx.ResponseWriter().Write([]byte("</form>"))
		
		return nil
	})
}

// SectionItem renders a section item with HTMX interactions
func SectionItem(sectionID, profileID, title, content string) templ.Component {
	return templ.ComponentFunc(func(ctx templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual HTMX component
		ctx.ResponseWriter().Write([]byte("<div id=\"section-" + sectionID + "\" class=\"section-item\">"))
		ctx.ResponseWriter().Write([]byte("<h3>" + title + "</h3>"))
		ctx.ResponseWriter().Write([]byte("<div>" + content + "</div>"))
		ctx.ResponseWriter().Write([]byte("<button hx-delete=\"/profiles/" + profileID + "/sections/" + sectionID + "\" hx-target=\"#section-" + sectionID + "\" hx-swap=\"outerHTML\">Delete</button>"))
		ctx.ResponseWriter().Write([]byte("</div>"))
		
		return nil
	})
}

// DeleteSectionButton renders a delete button for a section
func DeleteSectionButton(sectionID, profileID string) templ.Component {
	return templ.ComponentFunc(func(ctx templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual HTMX component
		ctx.ResponseWriter().Write([]byte("<button hx-delete=\"/profiles/" + profileID + "/sections/" + sectionID + "\" hx-target=\"#section-" + sectionID + "\" hx-swap=\"outerHTML\">Delete</button>"))
		
		return nil
	})
}