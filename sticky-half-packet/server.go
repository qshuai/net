package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic("listen error: " + err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	for {
		conn, err := l.Accept()
		if err != nil {
			cancel()
			panic(err)
		}

		go func(ctx context.Context) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Fprintf(os.Stderr, "tcp connection handler panic: %v", err)
					return
				}
			}()

			for {
				select {
				case <-ctx.Done():
					return
				default:
				}

				buf := make([]byte, 512)
				_, err = conn.Read(buf)
				if err != nil {
					if err == io.EOF {
						fmt.Printf("connection: %s closed\n", conn.RemoteAddr())
						return
					}

					panic(err)
				}

				// 模拟server处理延迟
				time.Sleep(time.Duration(rand.Int63n(10)) * time.Millisecond)

				fmt.Println(string(buf))
			}
		}(ctx)
	}
}

// 下面是一个tcp server端收到的一个包，同时含有半包和粘包的情况
// 这个吃瓜群众的编号是:960. 这个吃瓜群众的编号是:961. 这个吃瓜群众的编号是:962. 这个吃瓜群众的编号是:963. \
// 这个吃瓜群众的编号是:964. 这个吃瓜群众的编号是:965. 这个吃瓜群众的编号是:966. 这个吃瓜群众的编号是:967. \
// 这个吃瓜群众的编号是:968. 这个吃瓜群众的编号是:969. 这个吃瓜群众的编号是:970. 这个吃瓜群众的编号是:971. \
// 这个吃瓜群众的编号是:972. 这个吃瓜群众的编号是:973. 这个�
