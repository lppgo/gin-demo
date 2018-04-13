package helper

import (
	"demo/modules/common/base"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func GetValidator(c *gin.Context, validator *base.Validator, since string) *govalidator.Validator {
	if since != "" {
		validator.SetSence(since)
	}
	opts := govalidator.Options{
		Request:         c.Request,
		Rules:           validator.GetRules(),
		Messages:        validator.GetMessages(),
		RequiredDefault: true, //所有字段必须通过  true || false
	}
	return govalidator.New(opts)

}

func GetDefaultMsg(values map[string][]string) string {
	temp := ""
	for key, _ := range values {
		row, ok := values[key]
		if ok == true {
			temp = strings.TrimSpace(row[0])
			break
		}
	}
	return temp
}
