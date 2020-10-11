package main

import (
	. "fmt"

	_ "reflect" // 匿名导入，只会执行包的初始化函数

	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("log test")
	Println("fmt Println")
}

func init() {
	Println("init...")
}
