package main

import (
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		// 模拟网络延迟
		time.Sleep(time.Duration(rand.Int63n(2)) * time.Millisecond)

		_, err = conn.Write([]byte("这个吃瓜群众的编号是:" + strconv.Itoa(i+1) + ". "))
		if err != nil {
			panic(err)
		}
	}
}
