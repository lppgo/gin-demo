package validator

import (
	"github.com/gin-gonic/gin"
	"demo/modules/common/helper"
	"demo/modules/common/validator"
	"demo/modules/common/api/lang"
)

func valiData(c *gin.Context,validatorName string,sence string) (string,bool) {
	v := helper.GetValidator(c,validator.ValiFactory(validatorName),sence)
	err := v.Validate()
	if len(err) > 0{
		return lang.GetLangValiTip(c,helper.GetDefaultMsg(err)),false
	}
	return "",true
}
