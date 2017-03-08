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
		etext := fmt.Sprintf("Unknown opcode: %02X at address: %X", cpu.ram.read(cpu.PC), cpu.PC)
		//fmt.Printf("[% X ]\n", cpu.ram[:0x00FF])
		return errors.New(etext)
	}
	fmt.Printf("%04X|%02X|A:%02X|X:%02X|Y:%02X|P:%02X|SP:%02X\n", cpu.PC, i.opcode, cpu.A, cpu.X, cpu.Y, cpu.P, cpu.SP)
	cpu.numCycles += i.numCycles
	cpu.PC += i.size
	i.execute()
	return nil
}

/*
===============================================================================
				Addressing Modes (WIP)
===============================================================================
*/

//Implied Addressing: Not necessary to have an address returner?

//Accumulator Addressing: Not necessary to have an address returner?
func (cpu *CPU) accumulatorAddress() uint16 {
	return 0x000A //Temporary hack way to make A addressing work.
}

func (cpu *CPU) immediateAddress() uint16 {
	return cpu.PC - 1
}

func (cpu *CPU) zeroPageAddress() uint16 {
	hi := byte(0x00)
	lo := cpu.ram.read(cpu.immediateAddress())
	return binary.LittleEndian.Uint16([]byte{lo, hi})
}

func (cpu *CPU) zeroPageXAddress() uint16 {
	fmt.Println("ZERO X PAGE!!")
	d := cpu.ram.read(cpu.PC - 1)
	addr := byte(d + cpu.X)
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
	hi := cpu.ram.read(cpu.PC - 2)
	lo := cpu.ram.read(cpu.PC - 1)
	addr := binary.LittleEndian.Uint16([]byte{hi, lo})
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

func (cpu *CPU) indirectAddress() uint16 {
	var hi byte
	base := cpu.absoluteAddress()
	lo := cpu.ram.read(cpu.absoluteAddress())
	hi = cpu.ram.read(cpu.absoluteAddress() + 1)
	if base&0xFF > 0 {
		hi = cpu.ram.read(cpu.absoluteAddress() - 0xFF)
	}
	addr := binary.LittleEndian.Uint16([]byte{lo, hi})
	return addr
}

func (cpu *CPU) indexedIndirectAddress() uint16 {
	indirectLo := (cpu.ram.read(cpu.immediateAddress()) + cpu.X)
	indirectHi := byte(0x00)
	indirectAddr := binary.LittleEndian.Uint16([]byte{indirectLo, indirectHi})
	lo := indirectAddr
	if lo > 0xFF { // try to detect wrap around?
		lo = (lo - 0xFF)
	}
	hi := lo + 1
	if hi > 0xFF { // try to detect wrap around?
		hi = hi - (0xFF + 1)
	}
	addr := binary.LittleEndian.Uint16([]byte{cpu.ram.read(lo), cpu.ram.read(hi)})
	return addr
}

func (cpu *CPU) indirectIndexedAddress() uint16 {
	indirectLo := cpu.ram.read(cpu.immediateAddress())
	indirectHi := byte(0x00)
	lo := binary.LittleEndian.Uint16([]byte{indirectLo, indirectHi})
	if lo > 0xFF { // try to detect wrap around?
		lo = (lo - 0xFF)
	}
	hi := lo + 1
	if hi > 0xFF { // try to detect wrap around?
		hi = hi - (0xFF + 1)
	}
	addr := binary.LittleEndian.Uint16([]byte{cpu.ram.read(lo), cpu.ram.read(hi)})
	addr += uint16(cpu.Y)
	return addr
}

/*
===============================================================================
				Stack Operators
===============================================================================
*/
func (cpu *CPU) sPush(bytes ...byte) {
	for _, b := range bytes {
		addr := binary.LittleEndian.Uint16([]byte{cpu.SP, 0x01}) // stack
		cpu.ram.write(addr, b)
		cpu.SP--
	}
}

func (cpu *CPU) sPop() byte {
	cpu.SP++
	addr := binary.LittleEndian.Uint16([]byte{cpu.SP, 0x01}) //stack
	return cpu.ram.read(addr)
}

/*
===============================================================================
				Flag Setters (WIP)
===============================================================================
*/

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

//AND ... Logical AND performed between A register and contents of Memory (A&M)
func (cpu *CPU) AND(addr uint16) {
	M := cpu.ram.read(addr)
	cpu.A = (cpu.A & M)
	cpu.checkAndSetZeroFlag(cpu.A)
	cpu.checkAndSetNegativeFlag(cpu.A)
}

//ASL ... Arithmetic Shift Left
//A,Z,C,N = M*2 or M,Z,C,N = M*2
//This operation shifts all the bits of the accumulator or memory contents one bit left.
//Bit 0 is set to 0 and bit 7 is placed in the carry flag
//The effect of this operation is to multiply the memory contents by 2
//(ignoring 2's complement considerations), setting the carry if the result will not fit in 8 bits.
func (cpu *CPU) ASL(addr uint16) {
	var oval byte
	var nval byte
	if addr == uint16(0x000A) { //Work around to handle Accumulator addressing
		oval = cpu.A
		nval = oval << 1
		cpu.A = nval
	} else {
		oval = cpu.ram.read(addr)
		nval = oval << 1
		cpu.ram.write(addr, nval)
	}
	if hasBit(oval, 7) {
		cpu.P = setBit(cpu.P, 0)
	} else {
		cpu.P = clearBit(cpu.P, 0)
	}
	cpu.checkAndSetZeroFlag(nval)
	cpu.checkAndSetNegativeFlag(nval)
}

//BCC ... Branch if carry clear (If CPU.P.carry = false)
func (cpu *CPU) BCC(addr uint16) {
	if hasBit(cpu.P, 0) == false {
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

//BEQ ... Branch if equal (If CPU.P.zero = true)
func (cpu *CPU) BEQ(addr uint16) {
	if hasBit(cpu.P, 1) == true {
		displacement := uint16(cpu.ram.read(addr))
		cpu.PC += displacement
	}
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

//BMI ... Branch if minus
func (cpu *CPU) BMI(addr uint16) {
	if hasBit(cpu.P, 7) {
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

//BPL ... Branch if positive (If CPU.P.NegativeFlag = false, advance program counter)
func (cpu *CPU) BPL(addr uint16) {
	if hasBit(cpu.P, 7) == false {
		displacement := uint16(cpu.ram.read(addr))
		cpu.PC += displacement
	}
}

//BRK ... Force Interrupt
//The BRK instruction forces the generation of an interrupt request.
//The program counter and processor status are pushed on the stack then the IRQ
//interrupt vector at $FFFE/F is loaded into the PC and the break flag in the status set to one.
func (cpu *CPU) BRK() {
	//fmt.Println("BREAK ENCOUNTERED")
	//os.Exit(1)
}

//BVC ... Branch if Overflow Clear
func (cpu *CPU) BVC(addr uint16) {
	if hasBit(cpu.P, 6) == false {
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

//SEI ... Sets InterruptDisable flag on CPU status register
func (cpu *CPU) SEI() {
	cpu.P = setBit(cpu.P, 2)
}

//SEC ... Set Carry Flag
func (cpu *CPU) SEC() {
	cpu.P = setBit(cpu.P, 0)
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

//LSR ... Logical shift right
//A,C,Z,N = A/2 or M,C,Z,N = M/2
//Each of the bits in A or M is shift one place to the right.
//The bit that was in bit 0 is shifted into the carry flag. Bit 7 is set to zero
func (cpu *CPU) LSR(addr uint16) {
	var oval byte
	var nval byte
	if addr == uint16(0x000A) { //Work around to handle Accumulator addressing
		oval = cpu.A
		nval = oval >> 1
		cpu.A = nval
	} else {
		oval = cpu.ram.read(addr)
		nval = oval >> 1
		cpu.ram.write(addr, nval)
	}
	if hasBit(oval, 0) {
		cpu.P = setBit(cpu.P, 0)
	} else {
		cpu.P = clearBit(cpu.P, 0)
	}
	cpu.checkAndSetZeroFlag(nval)
	cpu.checkAndSetNegativeFlag(nval)
}

//NOP ... No Operation
func (cpu *CPU) NOP() {
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
	binary.BigEndian.PutUint16(bytes, cpu.PC-1)
	cpu.sPush(bytes...)
	cpu.PC = addr
}

//RTI ... Return from Interrupt
//The RTI instruction is used at the end of an interrupt processing routine.
//It pulls the processor flags from the stack followed by the program counter.
func (cpu *CPU) RTI() {
	cpu.PLP() //Pull processor status from stack
	lo := cpu.sPop()
	hi := cpu.sPop()
	cpu.PC = binary.LittleEndian.Uint16([]byte{lo, hi})
}

//RTS ...
func (cpu *CPU) RTS() {
	lo := cpu.sPop()
	hi := cpu.sPop()
	addr := binary.LittleEndian.Uint16([]byte{lo, hi})
	cpu.PC = addr + 1
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

//ROL ... Rotate Left
//Move each of the bits in either A or M one place to the left.
//Bit 0 is filled with the current value of the carry flag whilst
//the old bit 7 becomes the new carry flag value.
//TODO: THIS IS SUPER UGLY. FIX WITH PROPER BITSHIFTING EH?
func (cpu *CPU) ROL(addr uint16) {
	var oval byte
	var nval byte
	if addr == uint16(0x000A) {
		oval = cpu.A
		nval = oval << 1
		nval = clearBit(nval, 0)
		if hasBit(cpu.P, 0) {
			nval = setBit(nval, 0)
		}
		cpu.A = nval
	} else {
		oval = cpu.ram.read(addr)
		nval = oval << 1
		nval = clearBit(nval, 0)
		if hasBit(cpu.P, 0) {
			nval = setBit(nval, 0)
		}
		cpu.ram.write(addr, nval)
	}
	if hasBit(oval, 7) {
		cpu.P = setBit(cpu.P, 0)
	} else {
		cpu.P = clearBit(cpu.P, 0)
	}
	cpu.checkAndSetNegativeFlag(nval)
	cpu.checkAndSetZeroFlag(nval)
}

//ROR ... Rotate Right
//Move each of the bits in either A or M one place to the right.
//Bit 7 is filled with the current value of the carry flag whilst
//the old bit 0 becomes the new carry flag value.
//TODO: THIS IS SUPER UGLY. FIX WITH PROPER BITSHIFTING EH?
func (cpu *CPU) ROR(addr uint16) {
	var oval byte
	var nval byte
	if addr == uint16(0x000A) {
		oval = cpu.A
		nval = oval >> 1
		nval = clearBit(nval, 7)
		if hasBit(cpu.P, 0) {
			nval = setBit(nval, 7)
		}
		cpu.A = nval
	} else {
		oval = cpu.ram.read(addr)
		nval = oval >> 1
		nval = clearBit(nval, 7)
		if hasBit(cpu.P, 0) {
			nval = setBit(nval, 7)
		}
		cpu.ram.write(addr, nval)
	}
	if hasBit(oval, 0) {
		cpu.P = setBit(cpu.P, 0)
	} else {
		cpu.P = clearBit(cpu.P, 0)
	}
	cpu.checkAndSetNegativeFlag(nval)
	cpu.checkAndSetZeroFlag(nval)
}

//ORA ... Inclusive OR is performed between A register and contents of Memory (A|M)
func (cpu *CPU) ORA(addr uint16) {
	M := cpu.ram.read(addr)
	cpu.A = (cpu.A | M)
	cpu.checkAndSetZeroFlag(cpu.A)
	cpu.checkAndSetNegativeFlag(cpu.A)
}
