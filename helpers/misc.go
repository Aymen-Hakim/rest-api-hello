package helpers

import "github.com/gin-gonic/gin"

func JErr(c *gin.Context, err string) {
	c.JSON(400, gin.H{
		"err": err,
	})
}

func JOk(c *gin.Context, obj any) {
	c.JSON(200, obj)
}

func JDone(c *gin.Context) {
	JOk(c, gin.H{
		"done": true,
	})
}
