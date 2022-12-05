package main

import (
	"fmt"
	"os"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"github.com/mac-chaffee/generated-spoe/parser"
)

func parse(filename string) *parser.Spop {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	stream := kaitai.NewStream(f)
	spop := parser.NewSpop()
	if err = spop.Read(stream, nil, spop); err != nil {
		panic(err)
	}
	return spop
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <binary file>")
	}
	spop := parse(os.Args[1])
	fmt.Printf("Frame Length: %d\n", spop.FrameLen)
}
