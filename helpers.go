package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func readBinary(path string) []byte {
	bin, err := ioutil.ReadFile(path)
	check(err)

	return bin
}

func setBit(b byte, pos uint8) byte {
	b |= (1 << pos)
	return b
}

func clearBit(b byte, pos uint8) byte {
	mask := byte(^(1 << pos))
	b &= mask
	return b
}

func hasBit(b byte, pos uint8) bool {
	val := b & (1 << pos)
	return (val > 0)
}

func hexPrint(b byte) {
	fmt.Printf("%02X\n", b)
}
