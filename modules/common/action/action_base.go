package action

import (
	"demo/modules/common/api/redis"
)

type ActionBase struct {
	errMsg string
}

var delimiter = "_"

func (actionBase ActionBase) GetSendCode(account string, action string) string {
	_handler := redis.GetRedisCache()
	defer _handler.Close()
	result := _handler.Get(account + delimiter + action)
	if result == "" {
		return ""
	}
	return result
}

func (actionBase *ActionBase) ClearCode(account string, action string) {
	_handler := redis.GetRedisCache()
	defer _handler.Close()
	_handler.Expire(account+delimiter+action, 0)
}

func (actionBase *ActionBase) SetErrMsg(errMsg string) {
	actionBase.errMsg = errMsg
}

func (actionBase *ActionBase) GetErrMsg() string {
	return actionBase.errMsg
}
