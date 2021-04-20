/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/4/15
 * Time 下午5:41
 */

package ratelimit

import (
	"sync"
	"time"

	"imooc.com/ccmouse/learngo/limit/logger"
)

var limiter *Limiter

type Limiter struct {
	capacity int   //容量
	count    int   //当前量
	lastTime int64 //当前请求时间段
}

func GetInstance(num int) *Limiter {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}

	if limiter == nil {
		mutex.Lock()
		if limiter == nil {
			limiter = &Limiter{
				capacity: num,
				count:    0,
				lastTime: time.Now().Unix(),
			}
		}
		mutex.Unlock()
		wait.Wait()
	}
	return limiter
}

func Compare() bool {
	now := time.Now().Unix()
	if now > limiter.lastTime {
		limiter.count = 1
		limiter.lastTime = now
		return true
	}
	if now < limiter.lastTime {
		//异常情况，直接拒绝
		logger.Info.Println("time error")
		return false
	}
	if now == limiter.lastTime {
		if limiter.count < limiter.capacity {
			limiter.count++
			return true
		} else {
			return false
		}
	}
	return false
}
