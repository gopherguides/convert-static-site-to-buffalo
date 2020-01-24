package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func ContactHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("contact.html"))
}
