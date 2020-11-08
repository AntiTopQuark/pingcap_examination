package Process

import (
	"bufio"
	"log"
	"math"
	"os"
	"pingcap/Util"
	"strconv"
	"strings"
)

type Reducer struct {
	Split_Num int // 切分个数
	dic       string
}

var DefalutReducer = &Reducer{
	Split_Num: 16,
	dic:       "MapCollection",
}
var DefalutTrieReducer = &Reducer{
	Split_Num: 16,
	dic:       "TrieCollection",
}

func (receiver *Reducer) Process() (string, uint64) {

	log.Println("Reducer 任务开启")
	res_word := ""
	var res_index uint64 = math.MaxUint32
	for i := 0; i < receiver.Split_Num; i++ {
		// 读取数据
		log.Println("读取临时文件", i)
		tmpFileName := "tmp/_shuffle_file_" + strconv.Itoa(i)
		file, err := os.Open(tmpFileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		read_chan := Util.Reader(bufio.NewReader(file))

		// 构建Collection
		var collection Collection
		if receiver.dic == "MapCollection" {
			collection = &MapCollection{data: make(map[string]*Node)}
		} else if receiver.dic == "TrieCollection" {
			collection = &TrieCollection{data: Constructor()}
		}

		// 读取每一行数据，按照\t 进行切分，写入到dic中
		for v := range read_chan {
			lines := strings.Split(v, "\t")
			if len(lines) <= 1 {
				continue
			}
			word := lines[0]
			index, _ := strconv.ParseUint(lines[1], 10, 64)
			collection.Add(word, index)
		}

		// 获得该分区的数据结果
		tmp_word, tmp_index := collection.GetResult()
		if tmp_index < res_index {
			res_word = tmp_word
			res_index = tmp_index
		}
	}

	return res_word, res_index

}
