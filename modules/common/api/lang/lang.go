package lang

import (
	"github.com/gin-gonic/gin"
	"demo/utils/constant"
	"strings"
	"demo/lang"
	"demo/modules/common/helper"
)

func GetLang(c *gin.Context) string {
	headers := helper.GetRequestHeader(c)
	if len(headers) > 0 {
		lang, ok := headers[constant.PostHeaderLangField].(string)
		if ok == true {
			return strings.ToLower(lang)
		} else {
			return strings.ToLower(constant.DefaultLang)
		}
	} else {
		return strings.ToLower(constant.DefaultLang)
	}
}

func GetLangTip(c *gin.Context, moduleName string, tip string) string {
	return lang.GetInstance().GetTip(GetLang(c), moduleName, tip)
}

func GetLangCTip(c *gin.Context, tip string) string {
	return lang.GetInstance().GetTip(GetLang(c), "common", tip)
}

func GetLangValiTip(c *gin.Context, tip string) string {
	return lang.GetInstance().GetTip(GetLang(c), "validator", tip)
}
