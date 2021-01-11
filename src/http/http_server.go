package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/gin-gonic/gin"
)

var srv = &http.Server{
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}

var skipPaths = []string{
	"/api/mon/ping",
}

// Start http server
func Start() {
	r := gin.Default()
	ConfigRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080

}
// Shutdown http server
func Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("cannot shutdown http server:", err)
		os.Exit(2)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		fmt.Println("shutdown http server timeout of 5 seconds.")
	default:
		fmt.Println("http server stopped")
	}
}
