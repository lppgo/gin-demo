package validator

import (
	"github.com/gin-gonic/gin"
)

func ValiUserLogin(c *gin.Context, validatorName string) (string, bool) {
	return valiData(c, validatorName, "login")
}