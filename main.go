package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	userArgs := os.Args
	var (
		lenght = 7
		times  = 1
		prefix = ``
		sufix  = ``
		err    error
	)
	for pos, arg := range userArgs {
		switch arg {
		case `-l`:
			lenght, err = strconv.Atoi(userArgs[pos+1])
			if err != nil {
				fmt.Println(`Você deve usar apenas numeros no tamanho solicitado! - `, err)
			}
		case `-t`:
			times, err = strconv.Atoi(userArgs[pos+1])
			if err != nil {
				fmt.Println(`Você deve usar apenas numeros na quantidade solicitado! - `, err)
			}
		case `-pf`:
			prefix = userArgs[pos+1]
		case `-sf`:
			sufix = userArgs[pos+1]
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(times)
	for index := 0; index < times; index++ {
		go func() {
			defer wg.Done()
			fmt.Println(prefix + RandStringBytesMaskImprSrc(lenght) + sufix)
		}()
	}
	wg.Wait()
}

// Thanks: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
