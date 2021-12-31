package common

import (
	"crypto/md5"
	"encoding/hex"
)

//用于对数据进行加密等
func EncodePassword(pwd string) string{
	pwd += "kratosyyds"//第一层普通加密
	h := md5.New()
	h.Write([]byte(pwd))
	pwd = hex.EncodeToString(h.Sum(nil))
	pwd = "kratosddys" + pwd//第二层普通加密
	h.Write([]byte(pwd))
	return hex.EncodeToString(h.Sum(nil))//返回加密结果
}