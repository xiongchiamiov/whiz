package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/create", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/:link_id", func(c *gin.Context) {
		//name := c.Param("link_id")
		c.Redirect(307, "http://example.com/")
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
