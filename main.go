/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/4/15
 * Time 上午11:10
 */

package main

import (
	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/limit/Interceptor"
)

func main() {
	Router()
}

func Router() {

	router := gin.Default()

	router.Use(Interceptor.Limit())

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
