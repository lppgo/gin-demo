package models

import "time"

type User struct {
	ID              int       `xorm:"pk not null autoincr id"`
	Username        string    `xorm:"username"`
	Moble           string    `xorm:"moble"`
	Mobletime       int       `xorm:"mobletime"`
	Password        string    `xorm:"password"`
	Tpwdsetting     string    `xorm:"tpwdsetting"`
	Paypassword     string    `xorm:"paypassword"`
	Invit_1         string    `xorm:"invit_1"`
	Invit_2         string    `xorm:"invit_2"`
	Invit_3         string    `xorm:"invit_3"`
	IdcardStatus    int       `xorm:"idcard_status"`
	IdcardAuthNum   int       `xorm:"idcard_auth_num"`
	IdcardAuthTime  time.Time `xorm:"idcard_auth_time"`
	Truename        string    `xorm:"truename"`
	Idcard          string    `xorm:"idcard"`
	BankAuthNum     int       `xorm:"bank_auth_num"`
	BankAuthDay     int       `xorm:"bank_auth_day"`
	Logins          int       `xorm:"logins"`
	Ga              string    `xorm:"ga"`
	Addip           string    `xorm:"addip"`
	Addr            string    `xorm:"addr"`
	Sort            int       `xorm:"sort"`
	Addtime         int       `xorm:"addtime"`
	Endtime         int       `xorm:"endtime"`
	Status          int       `xorm:"status"`
	Vip             int       `xorm:"vip"`
	Email           string    `xorm:"email"`
	Alipay          string    `xorm:"alipay"`
	Invit           string    `xorm:"invit"`
	Token           string    `xorm:"token"`
	Awardid         int       `xorm:"awardid"`
	Awardstatus     int       `xorm:"awardstatus"`
	Awardname       string    `xorm:"awardname"`
	AwardNumAll     int       `xorm:"awardNumAll"`
	AwardNumToday   int       `xorm:"awardNumToday"`
	AwardTotalToday int       `xorm:"awardTotalToday"`
	Awardtime       int       `xorm:"awardtime"`
}

func (*User) TableName() string {
	return "user"
}
