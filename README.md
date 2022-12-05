# generated-spoe

This repo contains a parser for the Stream Processing Offload Protocol (SPOP), generated via kaitai.

Official protocol description: https://github.com/haproxy/haproxy/blob/v2.7.0/doc/SPOE.txt

## Usage

Right now, this repo only contains the raw generated parser. In the future, I'll add a nicer wrapper API around the generated parser (work in progress). Until then, you can perform rudimentary parsing with the generated parser directly:

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
2. In the root of this repo, run `go generate ./...`

Any changes to `grammar/*.ksy` requires regeneration. Otherwise, do not edit the generated code in the `parser` folder!

## Testing

```
go test ./...
```
