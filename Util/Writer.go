package Util

import (
	"bufio"
	"log"
)

func Writer(writer *bufio.Writer, in <-chan string) {
	for y := range in {
		_, err := writer.Write([]byte(y))
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
	}
}
