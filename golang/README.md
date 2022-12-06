# Generated SPOE Parser for Golang

## Usage

Right now, this library only contains the raw generated parser. In the future, I'll add a nicer wrapper API around the generated parser (work in progress). Until then, you can perform rudimentary parsing with the generated parser directly:

```go
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
```

Example:

```
$ go run main.go resources/notify_frame.bin
Frame Length: 883
```

## Generation

1. Follow these instructions to install kaitai: https://kaitai.io/#quick-start
2. Inside this folder, run `go generate ./...`

Any changes to `grammar/*.ksy` requires regeneration. Otherwise, do not edit the generated code in the `parser` folder!

## Testing

```
go test ./...
```
