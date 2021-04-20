/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/4/15
 * Time 下午3:33
 */

package Interceptor

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

func Limit2() gin.HandlerFunc {
	rl := ratelimit.New(1)
	return func(c *gin.Context) {
		rl.Take()
		c.Next()
	}
}
