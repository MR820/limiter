/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/4/15
 * Time 上午11:27
 */

package redis

import (
	"time"

	"github.com/go-redis/redis"
	"imooc.com/ccmouse/learngo/limit/logger"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "120.55.81.229:7890",
		Password: "123456",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		logger.Info.Println(err)
	}

	if pong != "PONG" {
		logger.Info.Println("客户端链接redis服务端失败")
	} else {
		logger.Info.Println("客户端成功连接redis服务端")
	}

}

func Set(key, val string) bool {
	_, err := client.Set(key, val, time.Second*10).Result()
	if err != nil {
		logger.Info.Println(err)
		return false
	}
	return true
}

func Get(key string) string {
	val, err := client.Get(key).Result()
	if err != nil {
		logger.Info.Println(err)
	}
	return val
}

func Exists(key string) bool {
	isExist, err := client.Exists(key).Result()
	if err != nil {
		logger.Info.Println(err)
	}
	if isExist == 0 {
		return false
	}
	return true
}

func HSet(hashTable, key, val string) bool {
	isSetSuccessful, err := client.HSet(hashTable, key, val).Result()

	if err != nil {
		logger.Info.Println(err)
		return false
	}
	expire(hashTable, 10)
	//存在返回false，不存在返回true
	return isSetSuccessful
}

func HGet(hashTable, key string) string {
	val, err := client.HGet(hashTable, key).Result()
	if err != nil {
		logger.Info.Println(err)
	}
	return val
}

func HExists(hashTable, key string) bool {
	isExists, err := client.HExists(hashTable, key).Result()
	if err != nil {
		logger.Info.Println(err)
		return false
	}
	return isExists
}

func expire(key string, expire int) {
	_, err := client.Expire(key, time.Duration(expire)*time.Second).Result()
	if err != nil {
		logger.Info.Println(err)
	}
}

func HMGet(hashTable string, fields []string) []interface{} {

	vals, err := client.HMGet(hashTable, fields...).Result()
	if err != nil {
		logger.Info.Println(err)
	}
	return vals
}

func MGet(fields []string) []interface{} {
	vals, err := client.MGet(fields...).Result()
	if err != nil {
		logger.Info.Println(err)
	}
	return vals
}
