package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r := gin.Default()
	r.POST("/post", func(c *gin.Context) {
		b := new(bytes.Buffer)
		b.ReadFrom(c.Request.Body)
		rcl := c.Request.Header.Get("Content-Length")
		rt := c.Request.Header.Get("Request-Timeout")

		fmt.Printf("Content Length: %s\n and the body\n%s", rcl, b.String())
		fmt.Printf("\n Request Timeout: %s", rt)
	})
	r.POST("/timeout", func(c *gin.Context) {
		fmt.Printf("Waiting for 80 seconds")
		time.Sleep(80 * time.Second)
	})

	r.Run(":" + port)
}
