package main

import (
	"fmt"
	"log"
	"pingcap/Process"
)

func main() {
	Process.DefalutMapper.Process("Data/input.in")
	word, index := Process.DefalutReducer.Process()
	if word == "" {
		log.Println("没有找到第一个未重复数据")
	} else {
		fmt.Println(word, index)
	}
}
