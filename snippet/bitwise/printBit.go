package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

func main() {
	printBit(uint64(18446744073709551615), 2)
	printBit(uint32(753321), 2)
}

// TODO:
// handle binary strings that are not multiples of 4
// handle hexadecimal output

func printBit(bits interface{}, base int) {

	var b uint64

	switch bits.(type) {
	case (uint64):
		b = bits.(uint64)
	case (uint32):
		b = uint64(bits.(uint32))
	case (uint16):
		b = uint64(bits.(uint16))
	case (uint8):
		b = uint64(bits.(uint8))
	default:
		fmt.Println("Unknown type for bit container")
		return
	}

	s := strconv.FormatUint(uint64(b), 2)

	var buffer bytes.Buffer

	i := 0
	for _, v := range s {
		step := math.Mod(float64(i), float64(4))
		buffer.WriteString(string(v))
		if step == 3 {
			buffer.WriteString(" ")
		}
		i++
	}
	fmt.Println(buffer.String())

}
