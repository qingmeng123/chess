package tools

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}
func ValidatePwd(pwd, salt, dbpwd string) bool {
	return Md5Encode(pwd+salt) == dbpwd
}
func MakeDbPwd(pwd, salt string) string {
	return Md5Encode(pwd + salt)
}
