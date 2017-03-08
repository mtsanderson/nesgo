package main

//Instruction ... Represents an instruction
type Instruction struct {
	Name      string
	opcode    byte
	size      uint16
	numCycles int
	//AddressingMode string
	execute func()
}

func (cpu *CPU) loadInstructions() {
	cpu.Instructions = make(map[byte]Instruction)

	//AAX (UNOFFICIAL)
	cpu.Instructions[0x87] = Instruction{
		Name:      "AAX",
		opcode:    0x87,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.AAX(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x97] = Instruction{
		Name:      "AAX",
		opcode:    0x97,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.AAX(cpu.zeroPageYAddress()) }}

	cpu.Instructions[0x83] = Instruction{
		Name:      "AAX",
		opcode:    0x83,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.AAX(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0x8F] = Instruction{
		Name:      "AAX",
		opcode:    0x8F,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.AAX(cpu.absoluteAddress()) }}

	//ADC
	cpu.Instructions[0x69] = Instruction{
		Name:      "ADC",
		opcode:    0x69,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.ADC(cpu.immediateAddress()) }}

	cpu.Instructions[0x65] = Instruction{
		Name:      "ADC",
		opcode:    0x65,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.ADC(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x75] = Instruction{
		Name:      "ADC",
		opcode:    0x75,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.ADC(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x6D] = Instruction{
		Name:      "ADC",
		opcode:    0x6D,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.ADC(cpu.absoluteAddress()) }}

	cpu.Instructions[0x7D] = Instruction{
		Name:      "ADC",
		opcode:    0x7D,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.ADC(cpu.absoluteXAddress()) }}

	cpu.Instructions[0x79] = Instruction{
		Name:      "ADC",
		opcode:    0x79,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.ADC(cpu.absoluteYAddress()) }}

	cpu.Instructions[0x61] = Instruction{
		Name:      "ADC",
		opcode:    0x61,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.ADC(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0x71] = Instruction{
		Name:      "ADC",
		opcode:    0x71,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.ADC(cpu.indirectIndexedAddress()) }}

	//AND
	cpu.Instructions[0x29] = Instruction{
		Name:      "AND",
		opcode:    0x29,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.AND(cpu.immediateAddress()) }}

	cpu.Instructions[0x25] = Instruction{
		Name:      "AND",
		opcode:    0x25,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.AND(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x35] = Instruction{
		Name:      "AND",
		opcode:    0x35,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.AND(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x2D] = Instruction{
		Name:      "AND",
		opcode:    0x2D,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.AND(cpu.absoluteAddress()) }}

	cpu.Instructions[0x3D] = Instruction{
		Name:      "AND",
		opcode:    0x3D,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.AND(cpu.absoluteXAddress()) }}

	cpu.Instructions[0x39] = Instruction{
		Name:      "AND",
		opcode:    0x39,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.AND(cpu.absoluteYAddress()) }}

	cpu.Instructions[0x21] = Instruction{
		Name:      "AND",
		opcode:    0x21,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.AND(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0x31] = Instruction{
		Name:      "AND",
		opcode:    0x31,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.AND(cpu.indirectIndexedAddress()) }}

	//ASL
	cpu.Instructions[0x0A] = Instruction{
		Name:      "ASL",
		opcode:    0x0A,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.ASL(cpu.accumulatorAddress()) }}

	cpu.Instructions[0x06] = Instruction{
		Name:      "ASL",
		opcode:    0x06,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.ASL(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x16] = Instruction{
		Name:      "ASL",
		opcode:    0x16,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.ASL(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x0E] = Instruction{
		Name:      "ASL",
		opcode:    0x0E,
		size:      3,
		numCycles: 6,
		execute:   func() { cpu.ASL(cpu.absoluteAddress()) }}

	cpu.Instructions[0x1E] = Instruction{
		Name:      "ASL",
		opcode:    0x1E,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.ASL(cpu.absoluteXAddress()) }}

	//ASO (UNOFFICIAL)
	cpu.Instructions[0x07] = Instruction{
		Name:      "ASO",
		opcode:    0x07,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.ASO(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x17] = Instruction{
		Name:      "ASO",
		opcode:    0x17,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.ASO(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x0F] = Instruction{
		Name:      "ASO",
		opcode:    0x0F,
		size:      3,
		numCycles: 6,
		execute:   func() { cpu.ASO(cpu.absoluteAddress()) }}

	cpu.Instructions[0x1F] = Instruction{
		Name:      "ASO",
		opcode:    0x1F,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.ASO(cpu.absoluteXAddress()) }}

	cpu.Instructions[0x1B] = Instruction{
		Name:      "ASO",
		opcode:    0x1B,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.ASO(cpu.absoluteYAddress()) }}

	cpu.Instructions[0x03] = Instruction{
		Name:      "ASO",
		opcode:    0x03,
		size:      2,
		numCycles: 8,
		execute:   func() { cpu.ASO(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0x13] = Instruction{
		Name:      "ASO",
		opcode:    0x13,
		size:      2,
		numCycles: 8,
		execute:   func() { cpu.ASO(cpu.indirectIndexedAddress()) }}

	//BCC
	cpu.Instructions[0x90] = Instruction{
		Name:      "BCC",
		opcode:    0x90,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.BCC(cpu.immediateAddress()) }}

	//BCS
	cpu.Instructions[0xB0] = Instruction{
		Name:      "BCS",
		opcode:    0xB0,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.BCS(cpu.immediateAddress()) }}

	//BEQ
	cpu.Instructions[0xF0] = Instruction{
		Name:      "BEQ",
		opcode:    0xF0,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.BEQ(cpu.immediateAddress()) }}

	//BIT
	cpu.Instructions[0x24] = Instruction{
		Name:      "BIT",
		opcode:    0x24,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.BIT(cpu.zeroPageAddress()) }}

	//BIT
	cpu.Instructions[0x2C] = Instruction{
		Name:      "BIT",
		opcode:    0x2C,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.BIT(cpu.absoluteAddress()) }}

	//BMI
	cpu.Instructions[0x30] = Instruction{
		Name:      "BMI",
		opcode:    0x30,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.BMI(cpu.immediateAddress()) }}

	//BNE
	cpu.Instructions[0xD0] = Instruction{
		Name:      "BNE",
		opcode:    0xD0,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.BNE(cpu.immediateAddress()) }}

	//BPL
	cpu.Instructions[0x10] = Instruction{
		Name:      "BPL",
		opcode:    0x10,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.BPL(cpu.immediateAddress()) }}

	//BRK
	cpu.Instructions[0x00] = Instruction{
		Name:      "BRK",
		opcode:    0x00,
		size:      1,
		numCycles: 7,
		execute:   func() { cpu.BRK() }}

	//BVC
	cpu.Instructions[0x50] = Instruction{
		Name:      "BVC",
		opcode:    0x50,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.BVC(cpu.immediateAddress()) }}

	//BVS
	cpu.Instructions[0x70] = Instruction{
		Name:      "BVS",
		opcode:    0x70,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.BVS(cpu.immediateAddress()) }}

	//CLC
	cpu.Instructions[0x18] = Instruction{
		Name:      "CLC",
		opcode:    0x18,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.CLC() }}

	//CLD
	cpu.Instructions[0xD8] = Instruction{
		Name:      "CLD",
		opcode:    0xD8,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.CLD() }}

	//CLI
	cpu.Instructions[0x58] = Instruction{
		Name:      "CLD",
		opcode:    0x58,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.CLI() }}

	//CLV
	cpu.Instructions[0xB8] = Instruction{
		Name:      "CLV",
		opcode:    0xB8,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.CLV() }}

	//CMP
	cpu.Instructions[0xC9] = Instruction{
		Name:      "CMP",
		opcode:    0xC9,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.CMP(cpu.immediateAddress()) }}

	cpu.Instructions[0xC5] = Instruction{
		Name:      "CMP",
		opcode:    0xC5,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.CMP(cpu.zeroPageAddress()) }}

	cpu.Instructions[0xD5] = Instruction{
		Name:      "CMP",
		opcode:    0xD5,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.CMP(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0xCD] = Instruction{
		Name:      "CMP",
		opcode:    0xCD,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.CMP(cpu.absoluteAddress()) }}

	cpu.Instructions[0xDD] = Instruction{
		Name:      "CMP",
		opcode:    0xDD,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.CMP(cpu.absoluteXAddress()) }}

	cpu.Instructions[0xD9] = Instruction{
		Name:      "CMP",
		opcode:    0xD9,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.CMP(cpu.absoluteYAddress()) }}

	cpu.Instructions[0xC1] = Instruction{
		Name:      "CMP",
		opcode:    0xC1,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.CMP(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0xD1] = Instruction{
		Name:      "CMP",
		opcode:    0xD1,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.CMP(cpu.indirectIndexedAddress()) }}

	//CPX
	cpu.Instructions[0xE0] = Instruction{
		Name:      "CPX",
		opcode:    0xE0,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.CPX(cpu.immediateAddress()) }}

	cpu.Instructions[0xE4] = Instruction{
		Name:      "CPX",
		opcode:    0xE4,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.CPX(cpu.zeroPageAddress()) }}

	cpu.Instructions[0xEC] = Instruction{
		Name:      "CPX",
		opcode:    0xEC,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.CPX(cpu.absoluteAddress()) }}

	//CPY
	cpu.Instructions[0xC0] = Instruction{
		Name:      "CPY",
		opcode:    0xC0,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.CPY(cpu.immediateAddress()) }}

	cpu.Instructions[0xC4] = Instruction{
		Name:      "CPY",
		opcode:    0xC4,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.CPY(cpu.zeroPageAddress()) }}

	cpu.Instructions[0xCC] = Instruction{
		Name:      "CPY",
		opcode:    0xCC,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.CPY(cpu.absoluteAddress()) }}

	//DCP (UNOFFICIAL)
	cpu.Instructions[0xC7] = Instruction{
		Name:      "DCP",
		opcode:    0xC7,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.DCP(cpu.zeroPageAddress()) }}

	cpu.Instructions[0xD7] = Instruction{
		Name:      "DCP",
		opcode:    0xD7,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.DCP(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0xCF] = Instruction{
		Name:      "DCP",
		opcode:    0xCF,
		size:      3,
		numCycles: 6,
		execute:   func() { cpu.DCP(cpu.absoluteAddress()) }}

	cpu.Instructions[0xDF] = Instruction{
		Name:      "DCP",
		opcode:    0xDF,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.DCP(cpu.absoluteXAddress()) }}

	cpu.Instructions[0xDB] = Instruction{
		Name:      "DCP",
		opcode:    0xDB,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.DCP(cpu.absoluteYAddress()) }}

	cpu.Instructions[0xC3] = Instruction{
		Name:      "DCP",
		opcode:    0xC3,
		size:      2,
		numCycles: 8,
		execute:   func() { cpu.DCP(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0xD3] = Instruction{
		Name:      "DCP",
		opcode:    0xD3,
		size:      2,
		numCycles: 8,
		execute:   func() { cpu.DCP(cpu.indirectIndexedAddress()) }}

	//DEC
	cpu.Instructions[0xC6] = Instruction{
		Name:      "DEC",
		opcode:    0xC6,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.DEC(cpu.zeroPageAddress()) }}

	cpu.Instructions[0xD6] = Instruction{
		Name:      "DEC",
		opcode:    0xD6,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.DEC(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0xCE] = Instruction{
		Name:      "DEC",
		opcode:    0xCE,
		size:      3,
		numCycles: 6,
		execute:   func() { cpu.DEC(cpu.absoluteAddress()) }}

	cpu.Instructions[0xDE] = Instruction{
		Name:      "DEC",
		opcode:    0xDE,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.DEC(cpu.absoluteXAddress()) }}

	//DEX
	cpu.Instructions[0xCA] = Instruction{
		Name:      "DEX",
		opcode:    0xCA,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.DEX() }}

	//DEY
	cpu.Instructions[0x88] = Instruction{
		Name:      "DEY",
		opcode:    0x88,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.DEY() }}

	//EOR
	cpu.Instructions[0x49] = Instruction{
		Name:      "EOR",
		opcode:    0x49,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.EOR(cpu.immediateAddress()) }}

	cpu.Instructions[0x45] = Instruction{
		Name:      "EOR",
		opcode:    0x45,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.EOR(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x55] = Instruction{
		Name:      "EOR",
		opcode:    0x55,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.EOR(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x4D] = Instruction{
		Name:      "EOR",
		opcode:    0x4D,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.EOR(cpu.absoluteAddress()) }}

	cpu.Instructions[0x5D] = Instruction{
		Name:      "EOR",
		opcode:    0x5D,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.EOR(cpu.absoluteXAddress()) }}

	cpu.Instructions[0x59] = Instruction{
		Name:      "EOR",
		opcode:    0x59,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.EOR(cpu.absoluteYAddress()) }}

	cpu.Instructions[0x41] = Instruction{
		Name:      "EOR",
		opcode:    0x41,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.EOR(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0x51] = Instruction{
		Name:      "EOR",
		opcode:    0x51,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.EOR(cpu.indirectIndexedAddress()) }}

	//INC
	cpu.Instructions[0xE6] = Instruction{
		Name:      "INC",
		opcode:    0xE6,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.INC(cpu.zeroPageAddress()) }}

	cpu.Instructions[0xF6] = Instruction{
		Name:      "INC",
		opcode:    0xF6,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.INC(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0xEE] = Instruction{
		Name:      "INC",
		opcode:    0xEE,
		size:      3,
		numCycles: 6,
		execute:   func() { cpu.INC(cpu.absoluteAddress()) }}

	cpu.Instructions[0xFE] = Instruction{
		Name:      "INC",
		opcode:    0xFE,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.INC(cpu.absoluteXAddress()) }}

	//ISC (UNOFFICIAL)
	cpu.Instructions[0xE7] = Instruction{
		Name:      "ISC",
		opcode:    0xE7,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.ISC(cpu.zeroPageAddress()) }}

	cpu.Instructions[0xF7] = Instruction{
		Name:      "ISC",
		opcode:    0xF7,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.ISC(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0xEF] = Instruction{
		Name:      "ISC",
		opcode:    0xEF,
		size:      3,
		numCycles: 6,
		execute:   func() { cpu.ISC(cpu.absoluteAddress()) }}

	cpu.Instructions[0xFF] = Instruction{
		Name:      "ISC",
		opcode:    0xFF,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.ISC(cpu.absoluteXAddress()) }}

	cpu.Instructions[0xFB] = Instruction{
		Name:      "ISC",
		opcode:    0xFB,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.ISC(cpu.absoluteYAddress()) }}

	cpu.Instructions[0xE3] = Instruction{
		Name:      "ISC",
		opcode:    0xE3,
		size:      2,
		numCycles: 8,
		execute:   func() { cpu.ISC(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0xF3] = Instruction{
		Name:      "ISC",
		opcode:    0xF3,
		size:      2,
		numCycles: 8,
		execute:   func() { cpu.ISC(cpu.indirectIndexedAddress()) }}

	//INX
	cpu.Instructions[0xE8] = Instruction{
		Name:      "DEX",
		opcode:    0xE8,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.INX() }}

	//INY
	cpu.Instructions[0xC8] = Instruction{
		Name:      "DEY",
		opcode:    0xC8,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.INY() }}

	//JMP
	cpu.Instructions[0x4C] = Instruction{
		Name:      "JMP",
		opcode:    0x4C,
		size:      3,
		numCycles: 3,
		execute:   func() { cpu.JMP(cpu.absoluteAddress()) }}

	cpu.Instructions[0x6C] = Instruction{
		Name:      "JMP",
		opcode:    0x6C,
		size:      3,
		numCycles: 5,
		execute:   func() { cpu.JMP(cpu.indirectAddress()) }}

	//JSR
	cpu.Instructions[0x20] = Instruction{
		Name:      "JSR",
		opcode:    0x20,
		size:      3,
		numCycles: 6,
		execute:   func() { cpu.JSR(cpu.absoluteAddress()) }}

	//LAX (UNOFFICIAL)
	cpu.Instructions[0xA7] = Instruction{
		Name:      "LAX",
		opcode:    0xA7,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.LAX(cpu.zeroPageAddress()) }}

	cpu.Instructions[0xB7] = Instruction{
		Name:      "LAX",
		opcode:    0xB7,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.LAX(cpu.zeroPageYAddress()) }}

	cpu.Instructions[0xAF] = Instruction{
		Name:      "LAX",
		opcode:    0xAF,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.LAX(cpu.absoluteAddress()) }}

	cpu.Instructions[0xBF] = Instruction{
		Name:      "LAX",
		opcode:    0xBF,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.LAX(cpu.absoluteYAddress()) }}

	cpu.Instructions[0xA3] = Instruction{
		Name:      "LAX",
		opcode:    0xA3,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.LAX(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0xB3] = Instruction{
		Name:      "LAX",
		opcode:    0xB3,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.LAX(cpu.indirectIndexedAddress()) }}

	//LDA
	cpu.Instructions[0xA9] = Instruction{
		Name:      "LDA",
		opcode:    0xA9,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.LDA(cpu.immediateAddress()) }}

	cpu.Instructions[0xA5] = Instruction{
		Name:      "LDA",
		opcode:    0xA5,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.LDA(cpu.zeroPageAddress()) }}

	cpu.Instructions[0xB5] = Instruction{
		Name:      "LDA",
		opcode:    0xB5,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.LDA(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0xAD] = Instruction{
		Name:      "LDA",
		opcode:    0xAD,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.LDA(cpu.absoluteAddress()) }}

	cpu.Instructions[0xBD] = Instruction{
		Name:      "LDA",
		opcode:    0xBD,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.LDA(cpu.absoluteXAddress()) }}

	cpu.Instructions[0xB9] = Instruction{
		Name:      "LDA",
		opcode:    0xB9,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.LDA(cpu.absoluteYAddress()) }}

	cpu.Instructions[0xA1] = Instruction{
		Name:      "LDA",
		opcode:    0xA1,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.LDA(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0xB1] = Instruction{
		Name:      "LDA",
		opcode:    0xB1,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.LDA(cpu.indirectIndexedAddress()) }}

	//LDX
	cpu.Instructions[0xA2] = Instruction{
		Name:      "LDX",
		opcode:    0xA2,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.LDX(cpu.immediateAddress()) }}

	cpu.Instructions[0xA6] = Instruction{
		Name:      "LDX",
		opcode:    0xA6,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.LDX(cpu.zeroPageAddress()) }}

	cpu.Instructions[0xB6] = Instruction{
		Name:      "LDX",
		opcode:    0xB6,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.LDX(cpu.zeroPageYAddress()) }}

	cpu.Instructions[0xAE] = Instruction{
		Name:      "LDX",
		opcode:    0xAE,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.LDX(cpu.absoluteAddress()) }}

	cpu.Instructions[0xBE] = Instruction{
		Name:      "LDX",
		opcode:    0xBE,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.LDX(cpu.absoluteYAddress()) }}

	//LDY
	cpu.Instructions[0xA0] = Instruction{
		Name:      "LDY",
		opcode:    0xA0,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.LDY(cpu.immediateAddress()) }}

	cpu.Instructions[0xA4] = Instruction{
		Name:      "LDY",
		opcode:    0xA4,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.LDY(cpu.zeroPageAddress()) }}

	cpu.Instructions[0xB4] = Instruction{
		Name:      "LDY",
		opcode:    0xB4,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.LDY(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0xAC] = Instruction{
		Name:      "LDY",
		opcode:    0xAC,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.LDY(cpu.absoluteAddress()) }}

	cpu.Instructions[0xBC] = Instruction{
		Name:      "LDY",
		opcode:    0xBC,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.LDY(cpu.absoluteXAddress()) }}

	//LSE (UNOFFICIAL)
	cpu.Instructions[0x47] = Instruction{
		Name:      "LSE",
		opcode:    0x47,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.LSE(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x57] = Instruction{
		Name:      "LSE",
		opcode:    0x57,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.LSE(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x4F] = Instruction{
		Name:      "LSE",
		opcode:    0x4F,
		size:      3,
		numCycles: 6,
		execute:   func() { cpu.LSE(cpu.absoluteAddress()) }}

	cpu.Instructions[0x5F] = Instruction{
		Name:      "LSE",
		opcode:    0x5F,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.LSE(cpu.absoluteXAddress()) }}

	cpu.Instructions[0x5B] = Instruction{
		Name:      "LSE",
		opcode:    0x5B,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.LSE(cpu.absoluteYAddress()) }}

	cpu.Instructions[0x43] = Instruction{
		Name:      "LSE",
		opcode:    0x43,
		size:      2,
		numCycles: 8,
		execute:   func() { cpu.LSE(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0x53] = Instruction{
		Name:      "LSE",
		opcode:    0x53,
		size:      2,
		numCycles: 8,
		execute:   func() { cpu.LSE(cpu.indirectIndexedAddress()) }}

	//LSR
	cpu.Instructions[0x4A] = Instruction{
		Name:      "LSR",
		opcode:    0x4A,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.LSR(cpu.accumulatorAddress()) }}

	cpu.Instructions[0x46] = Instruction{
		Name:      "LSR",
		opcode:    0x46,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.LSR(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x56] = Instruction{
		Name:      "LSR",
		opcode:    0x56,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.LSR(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x4E] = Instruction{
		Name:      "LSR",
		opcode:    0x4E,
		size:      3,
		numCycles: 6,
		execute:   func() { cpu.LSR(cpu.absoluteAddress()) }}

	cpu.Instructions[0x5E] = Instruction{
		Name:      "LSR",
		opcode:    0x5E,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.LSR(cpu.absoluteXAddress()) }}

	//NOP
	cpu.Instructions[0xEA] = Instruction{
		Name:      "NOP",
		opcode:    0xEA,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x1A] = Instruction{
		Name:      "NOP",
		opcode:    0x1A,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x3A] = Instruction{
		Name:      "NOP",
		opcode:    0x3A,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x5A] = Instruction{
		Name:      "NOP",
		opcode:    0x5A,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x7A] = Instruction{
		Name:      "NOP",
		opcode:    0x7A,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0xDA] = Instruction{
		Name:      "NOP",
		opcode:    0xDA,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0xFA] = Instruction{
		Name:      "NOP",
		opcode:    0xFA,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.NOP() }}

	//DOP (DOUBLE NOP) (UNOFFICIAL)
	cpu.Instructions[0x04] = Instruction{
		Name:      "NOP",
		opcode:    0x04,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x14] = Instruction{
		Name:      "NOP",
		opcode:    0x14,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x34] = Instruction{
		Name:      "NOP",
		opcode:    0x34,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x44] = Instruction{
		Name:      "NOP",
		opcode:    0x44,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x54] = Instruction{
		Name:      "NOP",
		opcode:    0x54,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x64] = Instruction{
		Name:      "NOP",
		opcode:    0x64,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x74] = Instruction{
		Name:      "NOP",
		opcode:    0x74,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x80] = Instruction{
		Name:      "NOP",
		opcode:    0x80,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x82] = Instruction{
		Name:      "NOP",
		opcode:    0x82,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x82] = Instruction{
		Name:      "NOP",
		opcode:    0x82,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x89] = Instruction{
		Name:      "NOP",
		opcode:    0x89,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0xD4] = Instruction{
		Name:      "NOP",
		opcode:    0xD4,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0xE2] = Instruction{
		Name:      "NOP",
		opcode:    0xE2,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0xF4] = Instruction{
		Name:      "NOP",
		opcode:    0xF4,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.NOP() }}

	//TOP (TRIPLE NOP) (UNOFFICIAL)
	cpu.Instructions[0x0C] = Instruction{
		Name:      "NOP",
		opcode:    0x0C,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x1C] = Instruction{
		Name:      "NOP",
		opcode:    0x1C,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x3C] = Instruction{
		Name:      "NOP",
		opcode:    0x3C,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x5C] = Instruction{
		Name:      "NOP",
		opcode:    0x5C,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0x7C] = Instruction{
		Name:      "NOP",
		opcode:    0x7C,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0xDC] = Instruction{
		Name:      "NOP",
		opcode:    0xDC,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.NOP() }}

	cpu.Instructions[0xFC] = Instruction{
		Name:      "NOP",
		opcode:    0xFC,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.NOP() }}

	//ORA
	cpu.Instructions[0x09] = Instruction{
		Name:      "ORA",
		opcode:    0x09,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.ORA(cpu.immediateAddress()) }}

	cpu.Instructions[0x05] = Instruction{
		Name:      "ORA",
		opcode:    0x05,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.ORA(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x15] = Instruction{
		Name:      "ORA",
		opcode:    0x15,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.ORA(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x0D] = Instruction{
		Name:      "ORA",
		opcode:    0x0D,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.ORA(cpu.absoluteAddress()) }}

	cpu.Instructions[0x1D] = Instruction{
		Name:      "ORA",
		opcode:    0x1D,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.ORA(cpu.absoluteXAddress()) }}

	cpu.Instructions[0x19] = Instruction{
		Name:      "ORA",
		opcode:    0x19,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.ORA(cpu.absoluteYAddress()) }}

	cpu.Instructions[0x01] = Instruction{
		Name:      "ORA",
		opcode:    0x01,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.ORA(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0x11] = Instruction{
		Name:      "ORA",
		opcode:    0x11,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.ORA(cpu.indirectIndexedAddress()) }}

	//PHA
	cpu.Instructions[0x48] = Instruction{
		Name:      "PHA",
		opcode:    0x48,
		size:      1,
		numCycles: 3,
		execute:   func() { cpu.PHA() }}

	//PHP
	cpu.Instructions[0x08] = Instruction{
		Name:      "PHP",
		opcode:    0x08,
		size:      1,
		numCycles: 3,
		execute:   func() { cpu.PHP() }}

	//PLA
	cpu.Instructions[0x68] = Instruction{
		Name:      "PLA",
		opcode:    0x68,
		size:      1,
		numCycles: 4,
		execute:   func() { cpu.PLA() }}

	//PLP
	cpu.Instructions[0x28] = Instruction{
		Name:      "PLP",
		opcode:    0x28,
		size:      1,
		numCycles: 4,
		execute:   func() { cpu.PLP() }}

	//RLA (UNOFFICIAL)
	cpu.Instructions[0x27] = Instruction{
		Name:      "RLA",
		opcode:    0x27,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.RLA(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x37] = Instruction{
		Name:      "RLA",
		opcode:    0x37,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.RLA(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x2F] = Instruction{
		Name:      "RLA",
		opcode:    0x2F,
		size:      3,
		numCycles: 6,
		execute:   func() { cpu.RLA(cpu.absoluteAddress()) }}

	cpu.Instructions[0x3F] = Instruction{
		Name:      "RLA",
		opcode:    0x3F,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.RLA(cpu.absoluteXAddress()) }}

	cpu.Instructions[0x3B] = Instruction{
		Name:      "RLA",
		opcode:    0x3B,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.RLA(cpu.absoluteYAddress()) }}

	cpu.Instructions[0x23] = Instruction{
		Name:      "RLA",
		opcode:    0x23,
		size:      2,
		numCycles: 8,
		execute:   func() { cpu.RLA(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0x33] = Instruction{
		Name:      "RLA",
		opcode:    0x33,
		size:      2,
		numCycles: 8,
		execute:   func() { cpu.RLA(cpu.indirectIndexedAddress()) }}

	//ROL
	cpu.Instructions[0x2A] = Instruction{
		Name:      "ROL",
		opcode:    0x2A,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.ROL(cpu.accumulatorAddress()) }}

	cpu.Instructions[0x26] = Instruction{
		Name:      "ROL",
		opcode:    0x26,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.ROL(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x36] = Instruction{
		Name:      "ROL",
		opcode:    0x36,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.ROL(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x2E] = Instruction{
		Name:      "ROL",
		opcode:    0x2E,
		size:      3,
		numCycles: 6,
		execute:   func() { cpu.ROL(cpu.absoluteAddress()) }}

	cpu.Instructions[0x3E] = Instruction{
		Name:      "ROL",
		opcode:    0x3E,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.ROL(cpu.absoluteXAddress()) }}

	//ROR
	cpu.Instructions[0x6A] = Instruction{
		Name:      "ROR",
		opcode:    0x6A,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.ROR(cpu.accumulatorAddress()) }}

	cpu.Instructions[0x66] = Instruction{
		Name:      "ROR",
		opcode:    0x66,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.ROR(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x76] = Instruction{
		Name:      "ROR",
		opcode:    0x76,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.ROR(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x6E] = Instruction{
		Name:      "ROR",
		opcode:    0x6E,
		size:      3,
		numCycles: 6,
		execute:   func() { cpu.ROR(cpu.absoluteAddress()) }}

	cpu.Instructions[0x7E] = Instruction{
		Name:      "ROR",
		opcode:    0x7E,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.ROR(cpu.absoluteXAddress()) }}

	//RRA (UNOFFICIAL)
	cpu.Instructions[0x67] = Instruction{
		Name:      "RRA",
		opcode:    0x67,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.RRA(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x77] = Instruction{
		Name:      "RRA",
		opcode:    0x77,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.RRA(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x6F] = Instruction{
		Name:      "RRA",
		opcode:    0x6F,
		size:      3,
		numCycles: 6,
		execute:   func() { cpu.RRA(cpu.absoluteAddress()) }}

	cpu.Instructions[0x7F] = Instruction{
		Name:      "RRA",
		opcode:    0x7F,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.RRA(cpu.absoluteXAddress()) }}

	cpu.Instructions[0x7B] = Instruction{
		Name:      "RRA",
		opcode:    0x7B,
		size:      3,
		numCycles: 7,
		execute:   func() { cpu.RRA(cpu.absoluteYAddress()) }}

	cpu.Instructions[0x63] = Instruction{
		Name:      "RRA",
		opcode:    0x63,
		size:      2,
		numCycles: 8,
		execute:   func() { cpu.RRA(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0x73] = Instruction{
		Name:      "RRA",
		opcode:    0x73,
		size:      2,
		numCycles: 8,
		execute:   func() { cpu.RRA(cpu.indirectIndexedAddress()) }}

	//RTI
	cpu.Instructions[0x40] = Instruction{
		Name:      "RTI",
		opcode:    0x40,
		size:      1,
		numCycles: 6,
		execute:   func() { cpu.RTI() }}

	//RTS
	cpu.Instructions[0x60] = Instruction{
		Name:      "RTS",
		opcode:    0x60,
		size:      1,
		numCycles: 6,
		execute:   func() { cpu.RTS() }}

	//SBC
	cpu.Instructions[0xE9] = Instruction{
		Name:      "SBC",
		opcode:    0xE9,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.SBC(cpu.immediateAddress()) }}

	cpu.Instructions[0xEB] = Instruction{
		Name:      "SBC",
		opcode:    0xEB,
		size:      2,
		numCycles: 2,
		execute:   func() { cpu.SBC(cpu.immediateAddress()) }}

	cpu.Instructions[0xE5] = Instruction{
		Name:      "SBC",
		opcode:    0xE5,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.SBC(cpu.zeroPageAddress()) }}

	cpu.Instructions[0xF5] = Instruction{
		Name:      "SBC",
		opcode:    0xF5,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.SBC(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0xED] = Instruction{
		Name:      "SBC",
		opcode:    0xED,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.SBC(cpu.absoluteAddress()) }}

	cpu.Instructions[0xFD] = Instruction{
		Name:      "SBC",
		opcode:    0xFD,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.SBC(cpu.absoluteXAddress()) }}

	cpu.Instructions[0xF9] = Instruction{
		Name:      "SBC",
		opcode:    0xF9,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.SBC(cpu.absoluteYAddress()) }}

	cpu.Instructions[0xE1] = Instruction{
		Name:      "SBC",
		opcode:    0xE1,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.SBC(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0xF1] = Instruction{
		Name:      "SBC",
		opcode:    0xF1,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.SBC(cpu.indirectIndexedAddress()) }}

	//SEC
	cpu.Instructions[0x38] = Instruction{
		Name:      "SEC",
		opcode:    0x38,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.SEC() }}

	//SED
	cpu.Instructions[0xF8] = Instruction{
		Name:      "SED",
		opcode:    0xF8,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.SED() }}

	//SEI
	cpu.Instructions[0x78] = Instruction{
		Name:      "SEI",
		opcode:    0x78,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.SEI() }}

	//STA
	cpu.Instructions[0x85] = Instruction{
		Name:      "STA",
		opcode:    0x85,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.STA(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x95] = Instruction{
		Name:      "STA",
		opcode:    0x95,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.STA(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x8D] = Instruction{
		Name:      "STA",
		opcode:    0x8D,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.STA(cpu.absoluteAddress()) }}

	cpu.Instructions[0x9D] = Instruction{
		Name:      "STA",
		opcode:    0x9D,
		size:      3,
		numCycles: 5,
		execute:   func() { cpu.STA(cpu.absoluteXAddress()) }}

	cpu.Instructions[0x99] = Instruction{
		Name:      "STA",
		opcode:    0x99,
		size:      3,
		numCycles: 5,
		execute:   func() { cpu.STA(cpu.absoluteYAddress()) }}

	cpu.Instructions[0x81] = Instruction{
		Name:      "STA",
		opcode:    0x81,
		size:      2,
		numCycles: 6,
		execute:   func() { cpu.STA(cpu.indexedIndirectAddress()) }}

	cpu.Instructions[0x91] = Instruction{
		Name:      "STA",
		opcode:    0x91,
		size:      2,
		numCycles: 5,
		execute:   func() { cpu.STA(cpu.indirectIndexedAddress()) }}

	//STX
	cpu.Instructions[0x86] = Instruction{
		Name:      "STX",
		opcode:    0x86,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.STX(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x96] = Instruction{
		Name:      "STX",
		opcode:    0x96,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.STX(cpu.zeroPageYAddress()) }}

	cpu.Instructions[0x8E] = Instruction{
		Name:      "STX",
		opcode:    0x8E,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.STX(cpu.absoluteAddress()) }}

	//STY
	cpu.Instructions[0x84] = Instruction{
		Name:      "STY",
		opcode:    0x84,
		size:      2,
		numCycles: 3,
		execute:   func() { cpu.STY(cpu.zeroPageAddress()) }}

	cpu.Instructions[0x94] = Instruction{
		Name:      "STY",
		opcode:    0x94,
		size:      2,
		numCycles: 4,
		execute:   func() { cpu.STY(cpu.zeroPageXAddress()) }}

	cpu.Instructions[0x8C] = Instruction{
		Name:      "STY",
		opcode:    0x8C,
		size:      3,
		numCycles: 4,
		execute:   func() { cpu.STY(cpu.absoluteAddress()) }}

	//TAX
	cpu.Instructions[0xAA] = Instruction{
		Name:      "TAX",
		opcode:    0xAA,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.TAX() }}

	//TAY
	cpu.Instructions[0xA8] = Instruction{
		Name:      "TAY",
		opcode:    0xA8,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.TAY() }}

	//TSX
	cpu.Instructions[0xBA] = Instruction{
		Name:      "TSX",
		opcode:    0xBA,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.TSX() }}

	//TXA
	cpu.Instructions[0x8A] = Instruction{
		Name:      "TXA",
		opcode:    0x8A,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.TXA() }}

	//TXS
	cpu.Instructions[0x9A] = Instruction{
		Name:      "TXS",
		opcode:    0x9A,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.TXS() }}

	//TYA
	cpu.Instructions[0x98] = Instruction{
		Name:      "TYA",
		opcode:    0x98,
		size:      1,
		numCycles: 2,
		execute:   func() { cpu.TYA() }}

}
