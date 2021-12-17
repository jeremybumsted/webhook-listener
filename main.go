package main

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/post", func(c *gin.Context) {
		b := new(bytes.Buffer)
		b.ReadFrom(c.Request.Body)
		res := b.String()
		fmt.Printf(res)
	})

	r.Run(":8080")
}
