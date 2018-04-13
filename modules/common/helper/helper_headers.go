package helper

import (
	"fmt"
	"demo/modules/common/api/redis"
	"demo/utils/json"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	userToken = "u:token:%d"
	userLogin = "u:login:%d"
	token     = "token"
	memberId  = "member_id"
)

func GetRequestHeader(c *gin.Context) map[string]interface{} {
	headers := strings.TrimSpace(c.DefaultPostForm("headers", ""))
	ResultVo := map[string]interface{}{}
	if headers == "" {
		return ResultVo
	}
	return json.StringToJson(headers)
}

func GetLoginUser(c *gin.Context) map[string]interface{} {
	hanlder := redis.GetRedisCache()
	defer hanlder.Close()
	headers := strings.TrimSpace(c.DefaultPostForm("headers", ""))
	ResultVo := map[string]interface{}{}
	if headers == "" {
		return ResultVo
	}
	UserInfo := json.StringToJson(headers)
	uId, _ := strconv.Atoi(UserInfo[memberId].(string))
	tKey := fmt.Sprintf(userToken, uId)
	uKey := fmt.Sprintf(userLogin, uId)
	ResultVo[memberId] = hanlder.GetInt(uKey)
	ResultVo[token] = hanlder.Get(tKey)
	return ResultVo
}

func GetUserHeader(userId int, tok string) map[string]interface{} {
	headers := map[string]interface{}{}
	headers[token] = tok
	headers[memberId] = strconv.Itoa(userId)
	return headers
}
