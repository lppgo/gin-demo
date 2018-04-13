package user

import (
	"demo/modules/common"
	"demo/modules/common/action"
	"demo/modules/common/api/lang"
	"demo/modules/common/api/mysql"
	"demo/modules/common/helper"
	"demo/modules/demo/models"
	"demo/utils/constant"
	"strings"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) (*models.User, interface{}, interface{}, bool) {
	password := strings.TrimSpace(c.PostForm("password"))
	username := strings.TrimSpace(c.PostForm("phone"))

	var conditionField string
	conditionField = "`moble` = ?"
	Handler := mysql.GetDefault()
	var User = models.User{}
	reply, _ := Handler.Cols("`id`,`password`,`username`,`moble`,`status`,`truename`").Where(conditionField, username).Get(&User)
	if !reply {
		return &User, lang.GetLangTip(c, "Index", "userNonExist"), "", false
	}
	if User.Status != constant.EnableStatus {
		return &User, lang.GetLangTip(c, "Index", "userStop"), "", false
	}
	if common.Get_uer_pass(password) != User.Password {
		return &User, lang.GetLangTip(c, "Index", "usernameOrPasswordError"), "", false
	}

	MemberAction := action.NewMemberAction()
	jTk, _ := MemberAction.Login(User)
	UserVo := map[string]interface{}{}
	rUser := map[string]interface{}{}
	UserVo["id"] = User.ID
	UserVo["username"] = User.Username
	UserVo["mobile"] = User.Moble
	UserVo["token"] = jTk
	rUser["user"] = UserVo

	return &User, rUser, helper.GetUserHeader(User.ID, jTk), true
}
