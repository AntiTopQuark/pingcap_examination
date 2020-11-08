package Util

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Reader(reader *bufio.Reader) <-chan string {
	out := make(chan string, 1024)
	go func() {
		for {
			line, err := reader.ReadString('\n')
			line = strings.Trim(line, "\n")
			out <- line
			if err != nil || err == io.EOF {
				if err != nil {
					fmt.Errorf("ReadString error")
				}
				break
			}
		}
		close(out)
	}()
	return out
}
