package trouter

import (
	"demo/middleware/allow"
	"demo/modules/demo/handlers/index"
	"github.com/gin-gonic/gin"
)

// Run is
func Run(r *gin.Engine) {
	indexRouteGroup := r.Group("/index")
	indexRouteGroup.Use(allow.Filter)
	{
		indexRouteGroup.POST("/login", index.Login)
	}
}
