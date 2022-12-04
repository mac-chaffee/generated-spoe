package main

import (
	"fmt"
)

func encodeVarint(b []byte, i int) (int, error) {
	if len(b) == 0 {
		return 0, fmt.Errorf("encode varint: insufficient space in buffer")
	}

	if i < 240 {
		b[0] = byte(i)
		return 1, nil
	}

	n := 0

	b[n] = byte(i) | 240
	n++
	i = (i - 240) >> 4
	for i >= 128 {
		if n > len(b)-1 {
			return 0, fmt.Errorf("encode varint: insufficient space in buffer")
		}

		b[n] = byte(i) | 128
		n++
		i = (i - 128) >> 7
	}

	if n > len(b)-1 {
		return 0, fmt.Errorf("encode varint: insufficient space in buffer")
	}

	b[n] = byte(i)
	n++

	return n, nil
}

func decodeVarint(b []byte) (int, int, error) {
	if len(b) == 0 {
		return 0, 0, fmt.Errorf("decode varint: unterminated sequence")
	}
	val := int(b[0])
	off := 1
	if val < 240 {
		return val, 1, nil
	}
	r := uint(4)
	for {
		if off > len(b)-1 {
			return 0, 0, fmt.Errorf("decode varint: unterminated sequence")
		}

		v := int(b[off])
		val += v << r
		off++
		r += 7

		if v < 128 {
			break
		}
	}
	return val, off, nil
}

func main() {
	var intmap map[int][]byte = map[int][]byte{
		0:   {0x00},
		1:   {0x01},
		239: {0xef},
		// 240: {0xff, 0xf0},
		662: {0xfd, 0x1a},
	}
	fmt.Println(decodeVarint(intmap[662]))
	buf := make([]byte, 4)
	_, err := encodeVarint(buf, 300000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x", buf)
}
