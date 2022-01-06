package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

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
		rb := b.String()
		rh := c.Request.Header.Get("Content-Length")

		fmt.Printf("Request Headers:\n" + rh + "Request Body:\n" + rb)
	})

	r.Run(":" + port)
}
