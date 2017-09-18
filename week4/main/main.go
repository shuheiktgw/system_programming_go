package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"strings"
	"fmt"
	"bufio"
)

func section() {
	reader := strings.NewReader("Example of io,SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}

func bin() {
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32

	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}

func bin2() {
	file, err := os.Open("./Lenna.png")
	if err != nil {
		panic(err)
	}

	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}

func dumpChunk(chunk io.Reader) {
	var length int32

	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader

	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)

		if err == io.EOF {
			break
		}

		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length) + 12))
		offset, _ = file.Seek(int64(length + 8), 1)
	}

	return chunks
}

var source = `1行目
2行目
3行目`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		fmt.Printf("%#v\n", line)
	}
}

