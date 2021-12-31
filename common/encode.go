package common

import (
	"crypto/md5"
	"crypto/sha256"
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
	s := sha256.New()
	s.Write([]byte(hex.EncodeToString(h.Sum(nil))))
	return hex.EncodeToString(s.Sum([]byte("kratos")))//最后一层用sha256加密
}