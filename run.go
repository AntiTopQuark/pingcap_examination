package main

import (
	"fmt"
	"github.com/arl/statsviz"
	"log"
	"net/http"
	"pingcap/Process"
	"time"
)

func startMonitor() {
	// Register statsviz handlers on the default serve mux.
	log.Println("Monitor Page is : http://localhost:6060/debug/statsviz/")
	statsviz.RegisterDefault()
	http.ListenAndServe(":6060", nil)
}
func main() {

	go startMonitor()

	start := time.Now()

	Process.DefalutMapper.Process("Data/tiny.in")
	word, index := Process.DefalutReducer.Process()
	end := time.Now()
	log.Printf("运行时长: %s \n", end.Sub(start).String())
	if word == "" {
		log.Println("没有找到第一个未重复数据")
	} else {
		fmt.Println("结果为：", word)
		fmt.Println("索引：", index) // 索引是以0开始的行数
	}

}
