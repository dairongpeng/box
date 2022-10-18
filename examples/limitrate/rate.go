package main

import (
	"flag"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

var (
	limit ratelimit.Limiter
	rps   = flag.Int("rps", 100, "request per second")
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[GIN] ")
	log.SetOutput(gin.DefaultWriter)
}

func leakBucket() gin.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limit.Take()
		log.Print(color.CyanString("%v", now.Sub(prev)))
		prev = now
	}
}

func ginRun(rps int) {
	// 每秒可以通过多少个？默认100个，也就是每个请求间隔10ms
	limit = ratelimit.New(rps)

	app := gin.Default()
	app.Use(leakBucket())

	app.GET("/rate", func(ctx *gin.Context) {
		ctx.JSON(200, "rate limiting test")
	})

	log.Printf(color.CyanString("Current Rate Limit: %v requests/s", rps))
	app.Run(":8080")
}

func main() {
	flag.Parse()
	ginRun(*rps)
}
