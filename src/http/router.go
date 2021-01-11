package http

import (
	"github.com/gin-gonic/gin"
	"godemo/src/service"
)

// Config routes
func ConfigRouter(r *gin.Engine) {
	r.Static("/pub", "./pub")
	r.StaticFile("/favicon.ico", "./pub/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})

	sys := r.Group("/api/tools/sys")
	{
		sys.GET("/ping", ping)
		sys.GET("/pid", pid)
		sys.GET("/addr", addr)
	}

	codeabout := r.Group("/api/tools/codeabout")
	{
		codeabout.GET("/encode", service.Encode)
	}

	//if config.Get().Logger.Level == "DEBUG" {
	//	pprof.Register(r, "/api/monapi/debug/pprof")
	//}
}
