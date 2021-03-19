package common

import (
	"crypto/md5"
	"encoding/hex"
)

//加密字符串
func MD5(str string) string {
	h := md5.New()
	str = str + "加盐加盐"
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 输出加密结

}
