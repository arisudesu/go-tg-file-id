package internal

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func RLEDecode(bs []byte) []byte {
	var res []byte
	var last []byte

	for _, cur := range bs {
		if string(last) == "\x00" {
			res = append(res, bytes.Repeat(last, int(cur))...)
			last = nil
		} else {
			res = append(res, last...)
			last = []byte{cur}
		}
	}
	res = append(res, last...)
	return res
}

func TLDecode(bs []byte) (dec, rem []byte, err error) {
	l, bs := bs[0], bs[1:]
	resto := 0

	if l > 254 {
		return nil, nil, errors.New("length too big")
	}
	if l == 254 {
		ll, bs := binary.LittleEndian.Uint32(bs[:3]), bs[3:]
		dec, bs = bs[:ll], bs[ll:]
		resto = posmod(int(-ll), 4)
	} else {
		dec, bs = bs[:l], bs[l:]
		resto = posmod(int(-(l + 1)), 4)
	}
	return dec, bs[resto:], nil
}

func posmod(a, b int) (modulo int) {
	modulo = a % b
	if modulo < 0 {
		modulo += abs(b)
	}
	return
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
