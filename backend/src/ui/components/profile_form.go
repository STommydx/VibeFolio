package components

import (
	\"github.com/a-h/templ\"
)

// ProfileForm renders a profile form component
func ProfileForm(profileID, name, title, summary string, isPublic bool) templ.Component {
	return templ.ComponentFunc(func(ctx templ.ComponentContext) error {
		// In a real implementation, this would render HTML using templ
		// For now, we'll just return a mock implementation
		
		// This is a placeholder for the actual templ component
		ctx.ResponseWriter().Write([]byte(\"<form class=\\\"profile-form\\\">\"))\n\t\tctx.ResponseWriter().Write([]byte(\"<input type=\\\"text\\\" name=\\\"name\\\" value=\\\"\" + name + \"\\\" placeholder=\\\"Name\\\">\"))\n\t\tctx.ResponseWriter().Write([]byte(\"<input type=\\\"text\\\" name=\\\"title\\\" value=\\\"\" + title + \"\\\" placeholder=\\\"Title\\\">\"))\n\t\tctx.ResponseWriter().Write([]byte(\"<textarea name=\\\"summary\\\" placeholder=\\\"Summary\\\">\" + summary + \"</textarea>\"))\n\t\tctx.ResponseWriter().Write([]byte(\"<label><input type=\\\"checkbox\\\" name=\\\"is_public\\\" \" + checked(isPublic) + \"> Public</label>\"))\n\t\tctx.ResponseWriter().Write([]byte(\"<button type=\\\"submit\\\">Save Profile</button>\"))\n\t\tctx.ResponseWriter().Write([]byte(\"</form>\"))\n\t\t\n\t\treturn nil\n\t})\n}\n\n// Helper function to generate checked attribute\nfunc checked(value bool) string {\n\tif value {\n\t\treturn \"checked\"\n\t}\n\treturn \"\"\n}\n