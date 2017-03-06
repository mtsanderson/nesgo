package main

import "flag"

func main() {
	romPath := flag.String("rom", "", "Path to ROM file")
	flag.Parse()

	rom := ROM{data: readBinary(*romPath)}
	nes := NES{rom: rom}
	nes.powerOn()
}
