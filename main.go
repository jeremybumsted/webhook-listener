package main

import (
	"bytes"
	"encoding/json"
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
		rb, err := json.MarshalIndent(b.String(), " ", " ")
		if err != nil {
			log.Fatal(err)
		}
		rcl := c.Request.Header.Get("Content-Length")
		//rct := c.Request.Header.Get("Content-Type")

		fmt.Printf("Content Length: %s\n and the body\n%d", rcl, rb)
	})

	r.Run(":" + port)
}
