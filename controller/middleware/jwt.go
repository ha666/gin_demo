package middleware

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ha666/gin_demo/initial/config"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"net/http"
	"strings"
)

// 中间件，检查用户token
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//region 获取jwt信息
		tokenVal := c.Request.Header.Get("Authorization")
		if tokenVal == "" {
			c.AbortWithStatusJSON(http.StatusOK, model.Response{
				Code:    "middleware.jwt.error",
				Message: "请求未携带token，无权限访问",
			})
			return
		}
		if !strings.HasPrefix(tokenVal, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusOK, model.Response{
				Code:    "middleware.jwt.error",
				Message: "请求未携带token，无权限访问",
			})
			return
		}
		if golibs.Length(tokenVal) < 128 {
			c.AbortWithStatusJSON(http.StatusOK, model.Response{
				Code:    "middleware.jwt.error",
				Message: "请求未携带超过128位的token参数，无权限访问",
			})
			return
		}

		tokenVal = golibs.SubString(tokenVal, golibs.Length("Bearer "), golibs.Length(tokenVal)-golibs.Length("Bearer "))
		//endregion

		//region 解析jwt信息
		var jwtInfo model.JwtInfo
		tokenInfo, err := jwt.Parse(tokenVal, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return []byte(""), errors.New("签名方法不正确")
			}
			arrs := strings.Split(token.Raw, ".")
			if len(arrs) == 3 {
				b64, err := base64.RawStdEncoding.DecodeString(arrs[1])
				if err != nil {
					return []byte(""), err
				}
				err = json.Unmarshal(b64, &jwtInfo)
				if err != nil {
					return []byte(""), err
				}
			}
			return []byte(config.Config.App.Jwt.Secret), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, model.Response{
				Code:    "middleware.jwt.error",
				Message: "解析token失败:" + err.Error(),
			})
			return
		}
		if !tokenInfo.Valid {
			c.AbortWithStatusJSON(http.StatusOK, model.Response{
				Code:    "middleware.jwt.error",
				Message: "jwt验证不通过",
			})
			return
		}
		//endregion

		//region 验证jwt的值
		if jwtInfo.Exp < golibs.Unix() {
			c.AbortWithStatusJSON(http.StatusOK,
				model.Response{
					Code:    "middleware.jwt.error",
					Message: "jwt信息已过期",
				})
			return
		}
		if jwtInfo.IP != c.ClientIP() {
			c.AbortWithStatusJSON(http.StatusOK,
				model.Response{
					Code:    "middleware.jwt.error",
					Message: "IP地址错误:" + c.ClientIP(),
				})
			return
		}
		if jwtInfo.UserId <= 0 {
			c.AbortWithStatusJSON(http.StatusOK,
				model.Response{
					Code:    "middleware.jwt.error",
					Message: "userId出错",
				})
			return
		}
		c.Set("userId", jwtInfo.UserId)
		c.Set("userName", jwtInfo.UserName)
		c.Next()
		//endregion

	}
}
