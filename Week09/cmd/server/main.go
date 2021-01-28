package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server(ctx)
}

func server(ctx context.Context) error {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalf("listen error: %v \n", err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			conn, err := listen.Accept()
			if err != nil {
				log.Printf("accept error: %v \n", err)
				continue
			}
			go handleConn(ctx, conn)
		}
	}
}

func handleConn(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	fmt.Println("start")
	ch := make(chan []byte)
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		return read(ctx, conn, ch)
	})

	g.Go(func() error {
		return write(ctx, conn, ch)
	})

	g.Wait()
	fmt.Println("end")
}

func read(ctx context.Context, conn net.Conn, ch chan []byte) error {
	r := bufio.NewReader(conn)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			line, _, err := r.ReadLine()
			if err != nil {
				close(ch)
				return err
			}
			ch <- line
			fmt.Printf("read message is: %s \n", line)
		}
	}
}

func write(ctx context.Context, conn net.Conn, ch chan []byte) error {
	w := bufio.NewWriter(conn)
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
			w.WriteString("hello ")
			w.Write(line)
			w.WriteString("\n")
			w.Flush()
			fmt.Println("write message success")
		}
	}
}