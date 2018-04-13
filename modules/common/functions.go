package common

import "demo/utils/md5"

func Get_uer_pass(pass string) string {
	return md5.Md5EnCode(pass)
}