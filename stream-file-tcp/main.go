package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

type FileServer struct{}

func (fs *FileServer) Start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go fs.handleConnection(conn)
	}

}

func (fs *FileServer) handleConnection(conn net.Conn) {
	buffer := new(bytes.Buffer)
	defer conn.Close()

	for {
		var size int64
		err := binary.Read(conn, binary.LittleEndian, &size)
		if err != nil {
			fmt.Println("Error reading size:", err)
			return
		}
		n, err := io.CopyN(buffer, conn, size)
		if err != nil {

			return
		}

		fmt.Println("Received data:", buffer.Bytes())
		fmt.Printf("Recevied %d bytes:\n", n)
	}
}

func sendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}
	defer conn.Close()
	binary.Write(conn, binary.LittleEndian, int64(size))
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))
	if err != nil {
		return err
	}
	fmt.Printf("Sent %d bytes\n", n)
	return nil
}

func main() {
	go func() {

		time.Sleep(4 * time.Second)
		err := sendFile(40000000)
		if err != nil {
			fmt.Println("Error sending file:", err)
		}
	}()
	Server := &FileServer{}
	Server.Start()

}
