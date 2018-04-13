package allow

import (
	"github.com/gin-gonic/gin"
)

func Filter(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
