package main

//Memory ... Memory interface
type Memory interface {
	read()
	write()
}

//RAM ...
type RAM [0xFFFF]byte

func (r *RAM) read(addr uint16) byte {
	return r[addr]
}

func (r *RAM) write(addr uint16, val byte) {
	r[addr] = val
}
