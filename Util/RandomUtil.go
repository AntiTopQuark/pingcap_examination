package Util

import (
	"math/rand"
	"time"
)

func GetRandomString(count int) chan string {
	out := make(chan string)
	go func() {
		for j := 0; j < count; j++ {
			str := "abcdefghijklmnopqrstuvwxyz"
			bytes := []byte(str)
			result := []byte{}
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			l := rand.Intn(32) + 1
			for i := 0; i < l; i++ {
				result = append(result, bytes[r.Intn(len(bytes))])
			}
			result = append(result, []byte("\n")...)
			out <- string(result)
		}
		close(out)
	}()
	return out
}
