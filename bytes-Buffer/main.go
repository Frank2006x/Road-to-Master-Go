package main

import (
	"bytes"
	"fmt"
	"io"
)

func WriteTo(w io.Writer,msg []byte) error{
	_,err := w.Write(msg)
	if err != nil {
		return fmt.Errorf("failed to write to writer: %w", err)
	}
	return nil

}


func main() {
	buf:=new(bytes.Buffer)
	err:=WriteTo(buf,[]byte("hello"))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Print(buf.String())
	
}