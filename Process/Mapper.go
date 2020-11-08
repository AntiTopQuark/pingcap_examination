package Process

import (
	"bufio"
	"fmt"
	"hash"
	"hash/adler32"
	"log"
	"os"
	"pingcap/Util"
	"strconv"
)

type Mapper struct {
	Split_Num int         // 切分个数
	Hash_Func hash.Hash32 // hash函数
}

var DefalutMapper = &Mapper{
	Split_Num: 16,
	Hash_Func: adler32.New(),
}

func (receiver *Mapper) Process(file_name string) {
	log.Println("Mapper 任务开启")

	// 读取文件 获得数据
	log.Println("读取源数据文件：", file_name)
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//创建临时文件
	log.Println("创建Shuffle的临时文件,开启Writer协程")
	splitFile := make([]*bufio.Writer, receiver.Split_Num)
	splitChan := make([]chan string, receiver.Split_Num)
	for i := 0; i < receiver.Split_Num; i++ {
		tmpFileName := "tmp/_shuffle_file_" + strconv.Itoa(i)
		tmpFile, err := os.Create(tmpFileName)
		if err != nil {
			panic(err)
		}
		defer tmpFile.Close()
		splitFile[i] = bufio.NewWriter(tmpFile)
		splitChan[i] = make(chan string, 1)

		go Util.Writer(splitFile[i], splitChan[i])
	}

	// 进行hash shuffle
	log.Println("读取数据，进行Hash Shuffle")
	read_chan := Util.Reader(bufio.NewReader(file))
	var word_idx uint64 = 0
	for v := range read_chan {
		hash_res, err := receiver.Hash_Func.Write([]byte(v))
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		hash_res %= receiver.Split_Num
		splitChan[hash_res] <- v + "\t" + strconv.FormatUint(word_idx, 10) + "\n"

		word_idx++
	}

	//关闭通道
	log.Println("Shuffle完成，关闭通道")
	for i := 0; i < receiver.Split_Num; i++ {
		close(splitChan[i])
	}

}
