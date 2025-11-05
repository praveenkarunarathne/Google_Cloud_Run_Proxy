package main

import (
	"io"
	"net"
	"os"
	"sync"
)
func main() {
        listenAddr := ":" + os.Getenv("PORT")
        targetAddr := os.Getenv("V2RAY_SERVER_IP") + ":80"
        ln, err := net.Listen("tcp", listenAddr)
        if err != nil {
                return
        }
        for {
                conn, err := ln.Accept()
                if err != nil {
                        continue
                }
                go handleConnection(conn, targetAddr)
        }
}
func handleConnection(src net.Conn, targetAddr string) {
        dst, err := net.Dial("tcp", targetAddr)
        if err != nil {
                src.Close()
		return
        }
		
	var wg sync.WaitGroup
	wg.Add(2)

        go func() {
                io.Copy(dst, src)
                wg.Done()
        }()

        go func() {
                io.Copy(src, dst)
                wg.Done()
        }()
		
        wg.Wait()
	src.Close()
	dst.Close()
}
