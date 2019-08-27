package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"net/http"
)

// 中间件，限制只能内网访问
func IntranetMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !golibs.IsIntranetIP(c.ClientIP()) {
			c.AbortWithStatusJSON(http.StatusOK, model.Response{
				Code:    "middleware.intranet.error",
				Message: "只能内网访问",
			})
			return
		}
		c.Next()
	}
}
