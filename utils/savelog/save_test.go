package savelog

import "testing"

func TestWriteErrorLog(t *testing.T) {
	data := map[string]string{}
	data["user"] = "12312323"
	data["password"] = "12312323"
	writeLog("fileName","errorMsg","备注信息",data)
}
