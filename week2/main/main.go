package main

import (
	"os"
	"bytes"
	"fmt"
	"net"
	"io"
)

func file() {
	file, err := os.Create("text.txt")
	if err != nil {
		panic(err)
	}

	file.Write([]byte("os.File example\n"))
	file.Close()
}

func screen() {
	os.Stdout.Write([]byte("os.Stdout example\n"))
}

func buffer() {
	var buffer bytes.Buffer

	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}

func connection() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}

func multiWriter() {
	file, err := os.Create("multiwriter.txt")
	if err!= nil {
		panic(err)
	}

	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}
