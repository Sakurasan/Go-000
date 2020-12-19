package handlers

import (
	"Go-000/Week04/model"
	"Go-000/Week04/pkg/config"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()
var groupApi *gin.RouterGroup

//in the same package init executes in file'name alphabet order
func init() {
	if config.GetBool("App.enable_cors") {
		enableCorsMiddleware()
	}
	if sp := config.GetString("App.static_path"); sp != "" {
		r.Use(static.Serve("/", static.LocalFile(sp, true)))
		if config.GetBool("App.enable_not_found") {
			r.NoRoute(func(c *gin.Context) {
				file := path.Join(sp, "index.html")
				c.File(file)
			})
		}
	}
	prefix := config.GetString("App.api_prefix")
	api := "api"
	if prefix != "" {
		api = fmt.Sprintf("%s/%s", api, prefix)
	}
	groupApi = r.Group(api)

	if config.GetString("GinbroApp.env") != "prod" {
		r.GET("/hello", func(c *gin.Context) {
			c.JSON(200, "world")
		})
	}

}

//NewSvr for graceful shutdown
func NewSvr() http.Handler {
	return r
}

//ServerRun start the gin server
func ServerRun() {
	addr := config.GetString("App.addr")
	if config.GetBool("App.enable_https") {
		log.Fatal(autotls.Run(r, addr))
	} else {
		r.Run(addr)
	}
}

//Close gin GinbroApp
func Close() {
	model.Close()
}
