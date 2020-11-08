package main

import (
	"bufio"
	"fmt"
	"os"
	"pingcap/Util"
	"time"
)

func main() {
	//num := 1
	//for num < 16 {
	//	fmt.Printf("############   num:%d    ###############\n",num)
	//
	//	start := time.Now()
	//	const filename = "Data/input.in"
	//	file, err := os.Create(filename)
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer file.Close()
	//	p := Util.GetRandomString(10000000)
	//	Util.Writer(bufio.NewWriter(file),p)
	//
	//	f,_ := os.OpenFile(filename,1,1)
	//	fi,_ := f.Stat()
	//	fmt.Printf("文件大小: %d \n",fi.Size())
	//	end := time.Now()
	//	fmt.Printf("运行时间: %s \n",end.Sub(start).String())
	//
	//	num ++
	//}

	start := time.Now()
	const filename = "Data/tiny.in"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//Tiny 10000
	//Large 100 * 10000000
	p := Util.GetRandomString(10000)
	Util.Writer(bufio.NewWriter(file), p)
	end := time.Now()
	fmt.Printf("运行时间: %s \n", end.Sub(start).String())

	f, _ := os.OpenFile(filename, 1, 1)
	fi, _ := f.Stat()
	fmt.Printf("文件大小: %d \n", fi.Size())

}
