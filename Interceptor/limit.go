/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/4/15
 * Time 上午11:22
 */

package Interceptor

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/limit/ratelimit"
)

//固定窗口
func Limit() gin.HandlerFunc {
	limiter := ratelimit.NewSliding(100*time.Millisecond, time.Second, 10)
	return func(c *gin.Context) {
		if limiter.IsLimited() {
			c.String(http.StatusInternalServerError, "rate limit,Drop")
			c.Abort()
			return
		}
		c.Next()
		return
	}
}
