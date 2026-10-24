package home_api

import (
	"github.com/gin-gonic/gin"

	"github.com/daqing/airway/lib/render"
	"github.com/daqing/iNeovim-website/app/views/home"
)

func IndexAction(c *gin.Context) {
	render.HTML(c, home.Index(home.English))
}

func ChineseIndexAction(c *gin.Context) {
	render.HTML(c, home.Index(home.Chinese))
}
