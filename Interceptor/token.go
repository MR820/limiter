/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/4/15
 * Time 下午5:06
 */

package Interceptor

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func Limit3() gin.HandlerFunc {
	bt := ratelimit.NewBucketWithQuantum(time.Second*1, 1, 1)
	return func(c *gin.Context) {
		if bt.Available() == 0 {
			c.String(http.StatusOK, "rate limit,Drop")
			c.Abort()
			return
		} else {
			bt.Take(1)
		}
		c.Next()
		return
	}
}
