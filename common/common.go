//公共功能
package common

import (
	"bluebell/pkg/jwt"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

//加密字符串
func MD5(str string) string {
	h := md5.New()
	str = str + "加盐加盐"
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 输出加密结

}

//从登陆的token信息中获取当前登陆的用户名
func GetTokenUserid(c *gin.Context) int64 {
	return c.GetInt64(jwt.UserID)
}
