package content

import (
	"github.com/gin-gonic/gin"
	"github.com/albrow/forms"
	"demo/utils/json"
)

func GetRequestContents(c *gin.Context) string {
	data,_ := forms.Parse(c.Request)
	jsonString,_ :=  json.ToJsonString(data.Values)
	return jsonString
}
