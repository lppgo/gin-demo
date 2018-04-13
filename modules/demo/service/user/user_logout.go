package user

import (
	"github.com/gin-gonic/gin"
	"demo/modules/common/action"
	"demo/modules/common/helper"
)

func Logout(c *gin.Context) (string, bool) {
	MemberAction := action.NewMemberAction()
	uId := helper.GetUId(c)
	if uId > 0 {
		MemberAction.Logout(uId)
		return "", true
	}
	return "", false
}
