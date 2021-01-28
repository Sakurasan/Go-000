package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	baseCtx, cancel := context.WithCancel(context.Background())

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		// 开始goroutine监听连接
		go handleConn(baseCtx, conn)
		select {
		case <-baseCtx.Done():
			fmt.Println("tcp listen routine stoped")
			return
		default:
			continue
		}

	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
	cancel()
	listen.Close()
}

func handleConn(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	cm := make(chan []byte)
	ctx, cancle := context.WithCancel(ctx)
	defer cancle()
	// 读写缓冲区
	// rd := bufio.NewReader(conn)
	// wr := bufio.NewWriter(conn)

	eg, _ := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return read(ctx, conn, cm)
	})
	eg.Go(func() error {
		return write(ctx, conn, cm)
	})
	eg.Wait()
}

func GO(ctx context.Context, proc func(ctx context.Context)) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Goroutine panic， err:", err)
			}
		}()
		proc(ctx)
	}()
}

func read(ctx context.Context, conn net.Conn, ch chan []byte) error {
	buffer := bufio.NewReader(conn)
	for {
		select {
		case <-ctx.Done():
			conn.Close()
			return ctx.Err()
		default:
			line, _, err := buffer.ReadLine()
			if err != nil {
				close(ch)
				return err
			}
			ch <- line
		}
	}
}

func write(ctx context.Context, conn net.Conn, ch chan []byte) error {
	buffer := bufio.NewWriter(conn)
	for {
		select {
		case <-ctx.Done():

			return ctx.Err()
		default:
			line, ok := <-ch
			if !ok {
				return nil
			}
			if len(line) <= 0 {
				continue
			}
			buffer.Write(line)
			buffer.WriteString("\n")
			buffer.Flush()
		}
	}
}
