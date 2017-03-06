package main

//NES ...
type NES struct {
	cpu CPU
	//ram     RAM
	rom ROM
}

func (nes *NES) powerOn() {
	// initialize stuff
	nes.rom.load()
	nes.cpu.init(nes.rom)

	// start program exectuion
	nes.run()
}

func (nes *NES) run() {
	// program loop
	for {
		nes.cpu.Step()
	}
}
