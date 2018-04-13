package action

type MobileAction struct {
	ActionBase
	currLang string
}

func MessSignBean(actionNum string) string {
	action := map[string]string{
		"1": "xx",    //xx
		"2": "xxxx",  //xxxx
		"3": "xxxxx", //xxxxx
	}
	if result, ok := action[actionNum]; ok == true {
		return result
	}
	return ""
}

func (mobileAction *MobileAction) CheckMobileCode(currCode string, mobile string, actionCode string) (string, bool) {
	actionSign := MessSignBean(actionCode)
	_currRem := mobileAction.GetSendCode(mobile, actionSign)
	if _currRem == "" {
		return mobileAction.getLangMsg("notExist"), false
	}
	if _currRem != currCode {
		return mobileAction.getLangMsg("error"), false
	}
	mobileAction.ClearCode(mobile, actionSign)
	return "", true
}

func (mobileAction *MobileAction) SetCurrLang(currLang string) {
	mobileAction.currLang = currLang
}

//TODO::待优化
func (mobileAction *MobileAction) getLangMsg(tipName string) string {
	cnLang := map[string]string{
		"notExist": "短信验证码不存在!",
		"error":    "短信验证码错误!",
	}
	enLang := map[string]string{
		"notExist": "SMS verification code does not exist !",
		"error":    "SMS Code Error!",
	}
	result := ""
	switch mobileAction.currLang {
	case "cn":
		result = cnLang[tipName]
	case "en":
		result = enLang[tipName]
	}
	return result
}
