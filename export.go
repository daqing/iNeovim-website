package main

import (
	"github.com/daqing/airway/cmd"
	"github.com/daqing/airway/lib/static"

	"github.com/daqing/iNeovim-website/app/views/home"
)

// Static export renders both localized home pages together with the
// committed frontend bundle.
func init() {
	cmd.SetStaticPages(
		static.Page{Slug: "/", Component: home.Index(home.English)},
		static.Page{Slug: "/zh/", Component: home.Index(home.Chinese)},
	)
}
