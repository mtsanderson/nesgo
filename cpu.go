package main

import (
	"encoding/binary"
	"errors"
	"fmt"
)

//CPU ...
type CPU struct {
	PC           uint16 //Program counter
	A            byte   //Accumulator register
	X            byte   //X register
	Y            byte   //Y register
	P            byte   //Status register
	SP           byte   //Stack pointer
	Instructions map[byte]Instruction
	rom          ROM // ROM
	ram          RAM //TODO: Make this a memory mapper of some sort
	numCycles    int
}

func (cpu *CPU) init(rom ROM) {
	cpu.PC = 0xC000
	cpu.A = 0
	cpu.X = 0
	cpu.Y = 0
	cpu.SP = 0xFD
	cpu.numCycles = 0
	cpu.rom = rom
	cpu.P = 0x24
	cpu.loadInstructions()

	fmt.Printf("Number of instructions implemented: %d\n", len(cpu.Instructions))

	//fmt.Printf("Instructions implemented: %d\n", len(cpu.Instructions))

	for i := 0; i < len(cpu.ram); i++ {
		cpu.ram[i] = 0x00
	}
	for i := 0; i < len(rom.prgROM); i++ {
		cpu.ram[(0xC000 + i)] = rom.prgROM[i]
	}
}

//Step ...
func (cpu *CPU) Step() {
	instructon := cpu.Instructions[cpu.ram.read(cpu.PC)]
	err := cpu.executeInstruction(instructon)
	check(err)
}

//executeInstruction ... Executes a CPU instructon
func (cpu *CPU) executeInstruction(i Instruction) error {
	_, exists := cpu.Instructions[i.opcode]
	if !exists {
		etext := fmt.Sprintf("Unknown opcode: %X at address: %X", cpu.ram.read(cpu.PC), cpu.PC)
		return errors.New(etext)
	}
	fmt.Printf("%X|%02X|A:%02X|X:%02X|Y:%02X|P:%02X|SP:%02X\n", cpu.PC, i.opcode, cpu.A, cpu.X, cpu.Y, cpu.P, cpu.SP)
	cpu.numCycles += i.numCycles
	cpu.PC += i.size
	i.execute()
	return nil
}

func (cpu *CPU) sPush(bytes ...byte) {
	for _, b := range bytes {
		high := byte(0x01)
		low := cpu.SP
		addr := binary.LittleEndian.Uint16([]byte{low, high}) // stack
		cpu.ram.write(addr, b)
		cpu.SP--
	}
}

func (cpu *CPU) sPop() byte {
	cpu.SP++
	high := byte(0x01)
	low := cpu.SP
	addr := binary.LittleEndian.Uint16([]byte{low, high}) //stack
	return cpu.ram.read(addr)
}

func (cpu *CPU) immediateAddress() uint16 {
	return cpu.PC - 1
}

func (cpu *CPU) zeroPageAddress() uint16 {
	addr := binary.LittleEndian.Uint16([]byte{cpu.ram.read(cpu.PC - 1), 0x00})
	return addr
}

func (cpu *CPU) zeroPageXAddress() uint16 {
	d := cpu.ram.read(cpu.PC - 1)
	addr := d + cpu.X
	if addr > 0xFF {
		fmt.Println("SHOULDA WRAPPED")
	}
	return uint16(addr)
}

func (cpu *CPU) zeroPageYAddress() uint16 {
	d := cpu.ram.read(cpu.PC - 1)
	addr := binary.LittleEndian.Uint16([]byte{d + cpu.Y, 0x00})
	return addr
}

func (cpu *CPU) absoluteAddress() uint16 {
	high := cpu.ram.read(cpu.PC - 2)
	low := cpu.ram.read(cpu.PC - 1)
	addr := binary.LittleEndian.Uint16([]byte{high, low})
	return addr
}

func (cpu *CPU) absoluteXAddress() uint16 {
	addr := cpu.absoluteAddress() + uint16(cpu.X)
	return addr
}

func (cpu *CPU) absoluteYAddress() uint16 {
	addr := cpu.absoluteAddress() + uint16(cpu.Y)
	return addr
}

func (cpu *CPU) relativeAddress() uint16 {
	return cpu.PC
}

func (cpu *CPU) checkAndSetZeroFlag(val byte) {
	if val == 0 {
		cpu.P = setBit(cpu.P, 1)
	} else {
		cpu.P = clearBit(cpu.P, 1)
	}
}

func (cpu *CPU) checkAndSetNegativeFlag(val byte) {
	if val&(1<<7) > 0 {
		cpu.P = setBit(cpu.P, 7)
	} else {
		cpu.P = clearBit(cpu.P, 7)
	}
}

func (cpu *CPU) checkAndSetOverflowFlag(val byte) {
	if val&(1<<6) > 0 {
		cpu.P = setBit(cpu.P, 6)
	} else {
		cpu.P = clearBit(cpu.P, 6)
	}
}

/*
===============================================================================
													INSTRUCTIONS BEGIN
===============================================================================
*/

//ADC ... Add with Carry
//A,Z,C,N = A+M+C
func (cpu *CPU) ADC(addr uint16) {
	A := cpu.A
	M := cpu.ram.read(addr)
	C := byte(0x00)
	if hasBit(cpu.P, 0) == true {
		C = 0x01
	}

	cpu.A += (M + C)
	cpu.checkAndSetZeroFlag(cpu.A)
	cpu.checkAndSetNegativeFlag(cpu.A)
	if int(A)+int(M)+int(C) > 0xFF {
		cpu.P = setBit(cpu.P, 0)
	} else {
		cpu.P = clearBit(cpu.P, 0)
	}
	if (A^M)&0x80 == 0 && (A^cpu.A)&0x80 != 0 {
		cpu.P = setBit(cpu.P, 6)
	} else {
		cpu.P = clearBit(cpu.P, 6)
	}
}

//TAX ... X = A
func (cpu *CPU) TAX() {
	cpu.X = cpu.A
	cpu.checkAndSetZeroFlag(cpu.X)
	cpu.checkAndSetNegativeFlag(cpu.X)
}

//TAY ... Y = A
func (cpu *CPU) TAY() {
	cpu.Y = cpu.A
	cpu.checkAndSetZeroFlag(cpu.Y)
	cpu.checkAndSetNegativeFlag(cpu.Y)
}

//TSX ... X = S
func (cpu *CPU) TSX() {
	cpu.X = cpu.SP
	cpu.checkAndSetZeroFlag(cpu.X)
	cpu.checkAndSetNegativeFlag(cpu.X)
}

//TXA ... A = X
func (cpu *CPU) TXA() {
	cpu.A = cpu.X
	cpu.checkAndSetZeroFlag(cpu.A)
	cpu.checkAndSetNegativeFlag(cpu.A)
}

//TXS ... S = X
func (cpu *CPU) TXS() {
	cpu.SP = cpu.X
}

//TYA ... A = Y
func (cpu *CPU) TYA() {
	cpu.A = cpu.Y
	cpu.checkAndSetZeroFlag(cpu.A)
	cpu.checkAndSetNegativeFlag(cpu.A)
}

//BIT ... Bit Test
func (cpu *CPU) BIT(addr uint16) {
	val := cpu.ram.read(addr)
	if (cpu.A & val) == 0 {
		cpu.P = setBit(cpu.P, 1)
	} else {
		cpu.P = clearBit(cpu.P, 1)
	}
	cpu.checkAndSetOverflowFlag(val)
	cpu.checkAndSetNegativeFlag(val)
}

//SEI ... Sets InterruptDisable flag on CPU status register
func (cpu *CPU) SEI() {
	cpu.P = setBit(cpu.P, 2)
}

//SEC ... Set Carry Flag
func (cpu *CPU) SEC() {
	cpu.P = setBit(cpu.P, 0)
}

//NOP ... No Operation
func (cpu *CPU) NOP() {
	//
}

//SED ... Sets decimal flag
func (cpu *CPU) SED() {
	cpu.P = setBit(cpu.P, 3)
}

//CLC ... Clears Carry Flag
func (cpu *CPU) CLC() {
	cpu.P = clearBit(cpu.P, 0)
}

//CLD ... Clears DecimalMode flag
func (cpu *CPU) CLD() {
	cpu.P = clearBit(cpu.P, 3)
}

//CLI ... Clears InterruptDisable flag
func (cpu *CPU) CLI() {
	cpu.P = clearBit(cpu.P, 2)
}

//CLV ... Clears Overflow Flag
func (cpu *CPU) CLV() {
	cpu.P = clearBit(cpu.P, 6)
}

//CMP ...
func (cpu *CPU) CMP(addr uint16) {
	M := cpu.ram.read(addr)
	res := (cpu.A - M)
	if cpu.A >= M {
		cpu.P = setBit(cpu.P, 0)
	} else {
		cpu.P = clearBit(cpu.P, 0)
	}
	if cpu.A == M {
		cpu.P = setBit(cpu.P, 1)
	} else {
		cpu.P = clearBit(cpu.P, 1)
	}
	cpu.checkAndSetNegativeFlag(res)
}

//CPX ... Compare X register -- Z,C,N = X-M
func (cpu *CPU) CPX(addr uint16) {
	M := cpu.ram.read(addr)
	res := (cpu.X - M)
	if cpu.X >= M {
		cpu.P = setBit(cpu.P, 0)
	} else {
		cpu.P = clearBit(cpu.P, 0)
	}
	if cpu.X == M {
		cpu.P = setBit(cpu.P, 1)
	} else {
		cpu.P = clearBit(cpu.P, 1)
	}
	cpu.checkAndSetNegativeFlag(res)
}

//CPY ... Compare Y register -- Z,C,N = Y-M
func (cpu *CPU) CPY(addr uint16) {
	M := cpu.ram.read(addr)
	res := (cpu.Y - M)
	if cpu.Y >= M {
		cpu.P = setBit(cpu.P, 0)
	} else {
		cpu.P = clearBit(cpu.P, 0)
	}
	if cpu.Y == M {
		cpu.P = setBit(cpu.P, 1)
	} else {
		cpu.P = clearBit(cpu.P, 1)
	}
	cpu.checkAndSetNegativeFlag(res)
}

//DEC ... Decrement memory -- M,Z,N = M-1
func (cpu *CPU) DEC(addr uint16) {
	oval := cpu.ram.read(addr)
	nval := oval - 1
	cpu.ram.write(addr, nval)
	cpu.checkAndSetZeroFlag(nval)
	cpu.checkAndSetNegativeFlag(nval)
}

//DEX ... Decrement X Register -- X,Z,N = X-1
func (cpu *CPU) DEX() {
	cpu.X--
	cpu.checkAndSetZeroFlag(cpu.X)
	cpu.checkAndSetNegativeFlag(cpu.X)
}

//DEY ... Decrement Y Register -- Y,Z,N = Y-1
func (cpu *CPU) DEY() {
	cpu.Y--
	cpu.checkAndSetZeroFlag(cpu.Y)
	cpu.checkAndSetNegativeFlag(cpu.Y)
}

//EOR ... Exclusing OR is performed between A register and contents of Memory
//A,Z,N = A^M
func (cpu *CPU) EOR(addr uint16) {
	M := cpu.ram.read(addr)
	cpu.A = (cpu.A ^ M)
	cpu.checkAndSetZeroFlag(cpu.A)
	cpu.checkAndSetNegativeFlag(cpu.A)
}

//INC ... Increment memory -- M,Z,N = M+1
func (cpu *CPU) INC(addr uint16) {
	oval := cpu.ram.read(addr)
	nval := oval + 1
	cpu.ram.write(addr, nval)
	cpu.checkAndSetZeroFlag(nval)
	cpu.checkAndSetNegativeFlag(nval)
}

//INX ... Increment X Register -- X,Z,N = X+1
func (cpu *CPU) INX() {
	cpu.X++
	cpu.checkAndSetZeroFlag(cpu.X)
	cpu.checkAndSetNegativeFlag(cpu.X)
}

//INY ... Increment Y Register -- Y,Z,N = Y+1
func (cpu *CPU) INY() {
	cpu.Y++
	cpu.checkAndSetZeroFlag(cpu.Y)
	cpu.checkAndSetNegativeFlag(cpu.Y)
}

//LDA ... Loads the byte at location, addr, into the A register
func (cpu *CPU) LDA(addr uint16) {
	val := cpu.ram.read(addr)
	cpu.A = val
	cpu.checkAndSetZeroFlag(cpu.A)
	cpu.checkAndSetNegativeFlag(cpu.A)
}

//LDX ... Loads the byte at location, addr, into the X register
func (cpu *CPU) LDX(addr uint16) {
	val := cpu.ram.read(addr)
	cpu.X = val
	cpu.checkAndSetZeroFlag(val)
	cpu.checkAndSetNegativeFlag(val)
}

//LDY ... Loads the byte at location, addr, into the Y register
func (cpu *CPU) LDY(addr uint16) {
	val := cpu.ram.read(addr)
	cpu.Y = val
	cpu.checkAndSetZeroFlag(cpu.Y)
	cpu.checkAndSetNegativeFlag(cpu.Y)
}

//STA ... M = A
func (cpu *CPU) STA(addr uint16) {
	cpu.ram.write(addr, cpu.A)
}

//STX ... M = X
func (cpu *CPU) STX(addr uint16) {
	cpu.ram.write(addr, cpu.X)
}

//STY ... M = Y
func (cpu *CPU) STY(addr uint16) {
	cpu.ram.write(addr, cpu.Y)
}

//JMP ... Moves Program Counter to address, addr
func (cpu *CPU) JMP(addr uint16) {
	cpu.PC = addr
}

//JSR -- The JSR instruction pushes the address (minus one) of
//the return point on to the stack and then sets
//the program counter to the target memory address.
func (cpu *CPU) JSR(addr uint16) {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, cpu.PC)
	cpu.sPush(bytes...)
	cpu.PC = addr
}

//RTS ...
func (cpu *CPU) RTS() {
	high := cpu.sPop()
	low := cpu.sPop()
	addr := binary.LittleEndian.Uint16([]byte{low, high})
	cpu.PC = addr
}

//SBC ... Subtract with Carry
//Subtracts the contents of a memory location to the accumulator together with the
//not of the carry bit. If overflow occurs the carry bit is clear
//A,Z,C,N = A-M-(1-C)
func (cpu *CPU) SBC(addr uint16) {
	A := cpu.A
	M := cpu.ram.read(addr)
	C := byte(0x00)
	if hasBit(cpu.P, 0) == true {
		C = 0x01
	}
	cpu.A = A - M - (1 - C)
	cpu.checkAndSetZeroFlag(cpu.A)
	cpu.checkAndSetNegativeFlag(cpu.A)
	if int(A)-int(M)-int(1-C) >= 0 {
		cpu.P = setBit(cpu.P, 0)
	} else {
		cpu.P = clearBit(cpu.P, 0)
	}
	if (A^M)&0x80 != 0 && (A^cpu.A)&0x80 != 0 {
		cpu.P = setBit(cpu.P, 6)
	} else {
		cpu.P = clearBit(cpu.P, 6)
	}
}

//BPL ... Branch if positive (If CPU.P.NegativeFlag = false, advance program counter)
func (cpu *CPU) BPL(addr uint16) {
	if hasBit(cpu.P, 7) == false {
		displacement := uint16(cpu.ram.read(addr))
		cpu.PC += displacement
	}
}

//BVS ... Branch if Overflow Set
func (cpu *CPU) BVS(addr uint16) {
	if hasBit(cpu.P, 6) == true {
		displacement := uint16(cpu.ram.read(addr))
		cpu.PC += displacement
	}
}

//BVC ... Branch if Overflow Clear
func (cpu *CPU) BVC(addr uint16) {
	if hasBit(cpu.P, 6) == false {
		displacement := uint16(cpu.ram.read(addr))
		cpu.PC += displacement
	}
}

//BMI ... Branch if minus
func (cpu *CPU) BMI(addr uint16) {
	if hasBit(cpu.P, 7) {
		displacement := uint16(cpu.ram.read(addr))
		cpu.PC += displacement
	}
}

//BEQ ... Branch if equal (If CPU.P.zero = true)
func (cpu *CPU) BEQ(addr uint16) {
	if hasBit(cpu.P, 1) == true {
		displacement := uint16(cpu.ram.read(addr))
		cpu.PC += displacement
	}
}

//BNE ... Branch if not equal (If CPU.P.zero = false)
func (cpu *CPU) BNE(addr uint16) {
	if hasBit(cpu.P, 1) == false {
		displacement := uint16(cpu.ram.read(addr))
		cpu.PC += displacement
	}
}

//BCS ... Branch if carry set (If CPU.P.carry = true)
func (cpu *CPU) BCS(addr uint16) {
	if hasBit(cpu.P, 0) == true {
		displacement := uint16(cpu.ram.read(addr))
		cpu.PC += displacement
	}
}

//BCC ... Branch if carry clear (If CPU.P.carry = false)
func (cpu *CPU) BCC(addr uint16) {
	if hasBit(cpu.P, 0) == false {
		displacement := uint16(cpu.ram.read(addr))
		cpu.PC += displacement
	}
}

//PHA ... Push Accumulator to Stack
func (cpu *CPU) PHA() {
	cpu.sPush(cpu.A)
}

//PHP ... Push Processor Status to Stack
func (cpu *CPU) PHP() {
	b := cpu.P
	b = setBit(b, 5)
	b = setBit(b, 4)
	cpu.sPush(b)
}

//PLA ... Pull an 8 bit value from stack and stores in the Accumulator (setting zero and negative)
func (cpu *CPU) PLA() {
	cpu.A = cpu.sPop()
	cpu.checkAndSetZeroFlag(cpu.A)
	cpu.checkAndSetNegativeFlag(cpu.A)
}

//PLP ... Pull an 8 bit value from stack and into the processor status register
func (cpu *CPU) PLP() {
	val := cpu.sPop()
	val = setBit(val, 5)
	val = clearBit(val, 4)
	cpu.P = val
}

//AND ... Logical AND performed between A register and contents of Memory (A&M)
func (cpu *CPU) AND(addr uint16) {
	M := cpu.ram.read(addr)
	cpu.A = (cpu.A & M)
	cpu.checkAndSetZeroFlag(cpu.A)
	cpu.checkAndSetNegativeFlag(cpu.A)
}

//ORA ... Inclusive OR is performed between A register and contents of Memory (A|M)
func (cpu *CPU) ORA(addr uint16) {
	M := cpu.ram.read(addr)
	cpu.A = (cpu.A | M)
	cpu.checkAndSetZeroFlag(cpu.A)
	cpu.checkAndSetNegativeFlag(cpu.A)
}
