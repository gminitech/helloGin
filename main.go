// main.go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin")
	})
	//r.Run() // listen and serve on 0.0.0.0:8080
	r.Run(":28080")
}
