package actions

import (
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/packd"
)

// HomeHandler is a default handler to serve up
// a home page.
func GalleryHandler(c buffalo.Context) error {
	commercial := []string{}
	assetsBox.Walk(func(nm string, f packd.File) error {
		if strings.HasPrefix(nm, "assets/images/gallery/commercial/medium/") {
			commercial = append(commercial, strings.TrimPrefix(nm, "assets/"))
		}
		return nil
	})
	c.Set("commercial", commercial)

	theater := []string{}
	assetsBox.Walk(func(nm string, f packd.File) error {
		if strings.HasPrefix(nm, "assets/images/gallery/theater/medium/") {
			theater = append(theater, strings.TrimPrefix(nm, "assets/"))
		}
		return nil
	})
	c.Set("theater", theater)

	prewire := []string{}
	assetsBox.Walk(func(nm string, f packd.File) error {
		if strings.HasPrefix(nm, "assets/images/gallery/prewire/medium/") {
			prewire = append(prewire, strings.TrimPrefix(nm, "assets/"))
		}
		return nil
	})
	c.Set("prewire", prewire)

	return c.Render(200, r.HTML("gallery.html"))
}
