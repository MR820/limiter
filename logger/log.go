/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/3/25
 * Time 9:51 上午
 */

package logger

import (
	"io"
	"log"
	"os"
)

var Info *log.Logger

func init() {
	infoFile, err := os.OpenFile("./info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("打开日志文件失败", err)
	}

	Info = log.New(os.Stdout, "Info", log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(io.MultiWriter(os.Stderr, infoFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
}
