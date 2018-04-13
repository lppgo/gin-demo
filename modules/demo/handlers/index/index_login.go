package index

import (
	"demo/modules/common/api/lang"
	"demo/modules/demo/service/user"
	"demo/modules/demo/service/validator"
	"demo/utils/respo"
	"github.com/gin-gonic/gin"
)

// Login is
func Login(c *gin.Context) {
	if err, ok := validator.ValiUserLogin(c, "UserValidator"); ok == false {
		respo.HttpHErr(c, err, nil, nil, "")
		return
	}
	userInfo, result, headers, ok := user.Login(c)
	if ok == false {
		respo.HttpHErr(c, result.(string), nil, nil, "")
		return
	}

	respo.HttpHSucc(c, lang.GetLangTip(c, "Index", "loginSuccess"), headers, userInfo, "")
	return
}
