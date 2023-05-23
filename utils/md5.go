package utils

import (
	"crypto/md5"
	"encoding/hex"
	"im/log"
)

func Md5Encode(data string) string {	
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

func MakePwd(rawPwd, salt string) string {
	return Md5Encode(rawPwd+salt)
}

func CheckPwd(rawPwd, salt, pwd string) bool {
	log.Info.Println(rawPwd, salt, pwd)
	return MakePwd(rawPwd, salt) == pwd
}
