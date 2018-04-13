package validator

import (
	"github.com/gin-gonic/gin"
	"strings"
	"demo/modules/common/action"
	"demo/modules/common/api/lang"
)

func ValiSms(c *gin.Context, action string) (string, bool) {
	currCode := strings.TrimSpace(c.PostForm("smsCode"))
	mobile := strings.TrimSpace(c.PostForm("phone"))
	return GetMobileCheckHandler(c).CheckMobileCode(currCode, mobile, action)
}

func GetMobileCheckHandler(c *gin.Context) *action.MobileAction {
	mobileHandler := &action.MobileAction{}
	mobileHandler.SetCurrLang(lang.GetLang(c))
	return mobileHandler
}

func ValiSmsForm(c *gin.Context, validatorName string) (string, bool) {
	return valiData(c, validatorName, "mobileCode")
}
