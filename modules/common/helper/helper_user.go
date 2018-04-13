package helper

import (
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUId(c *gin.Context) int {
	var userInfo = GetLoginUser(c)
	if len(userInfo) > 0 {
		if reflect.TypeOf(userInfo[memberId]).String() == "string" {
			uId, _ := strconv.Atoi(userInfo[memberId].(string))
			return uId
		}
		return userInfo[memberId].(int)
	} else {
		return 0
	}
}
