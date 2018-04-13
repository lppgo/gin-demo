package action

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"demo/modules/common/api/redis"
	"demo/modules/demo/models"
	"strconv"

	"github.com/satori/go.uuid"
)

type MemberAction struct {
}

var (
	userToken    = "u:token:%d"
	userLogin    = "u:login:%d"
	regToken     = "u:register:%d"
	loginTime    = 3600
	registerTime = 60 * 20
)

func NewMemberAction() *MemberAction {
	var memberAction = new(MemberAction)
	return memberAction
}

func (MemberAction) Login(user models.User) (string, bool) {
	hanlder := redis.GetRedisCache()
	defer hanlder.Close()
	userId := user.ID
	_currToken := getUserToken(user)
	tKey := fmt.Sprintf(userToken, userId)
	uKey := fmt.Sprintf(userLogin, userId)
	hanlder.Set(tKey, _currToken, loginTime)
	hanlder.Set(uKey, userId, loginTime)
	return _currToken, true
}

func getUserToken(user models.User) string {
	userId := strconv.Itoa(user.ID)
	return sign(userId + user.Password)
}

func sign(sign string) string {
	Md5 := md5.New()
	Md5.Write([]byte(sign))
	return hex.EncodeToString(Md5.Sum(nil))
}

func CreateRegisterKey(uId int) string {
	handler := redis.GetRedisCache()
	defer handler.Close()
	token := uuid.NewV4().String()
	rTkey := fmt.Sprintf(regToken, uId)
	handler.Set(rTkey, token, registerTime)
	return token
}

func (MemberAction) Logout(userId int) {
	hanlder := redis.GetRedisCache()
	defer hanlder.Close()
	tKey := fmt.Sprintf(userToken, userId)
	uKey := fmt.Sprintf(userLogin, userId)
	hanlder.Expire(tKey, -1)
	hanlder.Expire(uKey, -1)
}

func CheckRegisterToken(uId int, rToken string) bool {
	handler := redis.GetRedisCache()
	defer handler.Close()
	rTkey := fmt.Sprintf(regToken, uId)
	currToken := handler.Get(rTkey)
	if currToken == rToken {
		return true
	}
	return false
}
