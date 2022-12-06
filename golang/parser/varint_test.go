package parser

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

// Map between the decimal value and the raw varint value
var intmap map[int][]byte = map[int][]byte{
	0:         {0x00},
	1:         {0x01},
	239:       {0xef},
	240:       {0xf0, 0x00},
	669:       {0xfd, 0x1a},
	20000:     {0xf0, 0xd3, 0x08},
	300000:    {0xf0, 0xaf, 0x91, 0x00},
	123456789: {0xf5, 0xc2, 0xf8, 0xd5, 0x02},
}

// Makes sure that the varint sub-grammar can parse numbers correctly
func TestVarInt(t *testing.T) {
	for expected, raw := range intmap {
		varint := NewVarint()
		buf := bytes.NewReader(raw)
		stream := kaitai.NewStream(buf)
		if err := varint.Read(stream, nil, varint); err != nil {
			panic(err)
		}

		decoded_value, err := varint.Value()
		assert.NoError(t, err)
		assert.Equal(t, expected, decoded_value)
	}

}
