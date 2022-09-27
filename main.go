package main

import (
	"github.com/gin-gonic/gin"
	"restapi.com/hello/queue"
)

func main() {
	q := queue.New()
	r := gin.Default()
	r.PUT("/cmd/:cmd", func(c *gin.Context) {
		q.Enqueue(c.Param("cmd"))
		c.JSON(200, gin.H{
			"done": "true",
		})
	})
	r.GET("/cmd", func(c *gin.Context) {
		cmd, err := q.Dequeue()
		if err != nil {
			cmd = "no command at the moment"
		}
		c.JSON(200, gin.H{
			"cmd": cmd,
		})
	})
	r.Run("127.0.0.1:8080")
}
