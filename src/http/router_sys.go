package http

import (
	"os"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.String(200, "pong")
}

func pid(c *gin.Context) {
	c.String(200, "%d", os.Getpid())
}

func addr(c *gin.Context) {
	c.String(200, c.Request.RemoteAddr)
}
