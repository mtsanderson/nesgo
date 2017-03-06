package main

const headerSize int = 16
const kbSize int = 1024

//ROM ...
type ROM struct {
	header  []byte
	prgSize int
	chrSize int
	trainer bool
	prgROM  []byte
	chrROM  []byte
	data    []byte
}

func (rom *ROM) load() {
	// initialize ROM
	rom.header = rom.data[:headerSize]
	rom.prgSize = int(rom.header[4])
	rom.chrSize = int(rom.header[5])
	//TODO: implement trainer detection
	//rom.prgROM = rom.data[headerSize:((kbSize * headerSize * rom.prgSize) + 16)]
	rom.prgROM = rom.data[headerSize:0x4000]
	rom.chrROM = rom.data[16+len(rom.prgROM):]
}
