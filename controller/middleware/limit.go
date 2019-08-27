package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ha666/gin_demo/model"
	"github.com/yangwenmai/ratelimit/simpleratelimit"
	"net/http"
	"time"
)

var (
	rl = simpleratelimit.New(100000, time.Minute)
)

// 中间件，用令牌桶限制请求频率
func LimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if rl.Limit() {
			c.AbortWithStatusJSON(
				http.StatusOK,
				model.Response{
					Code:    "middleware.limit",
					Message: "请求频率太高",
				},
			)
			return
		}
		c.Next()
	}
}
