package vm

const (
	Load  = 0x01
	Store = 0x02
	Add   = 0x03
	Sub   = 0x04
	Halt  = 0xff
)

// Stretch goals
const (
	Addi = 0x05
	Subi = 0x06
	Jump = 0x07
	Beqz = 0x08
)

// Given a 256 byte array of "memory", run the stored program
// to completion, modifying the data in place to reflect the result
//
// The memory format is:
//
// 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f ... ff
// __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ ... __
// ^==DATA===============^ ^==INSTRUCTIONS==============^
//
func compute(memory []byte) {

	registers := [3]byte{8, 0, 0} // PC, R1 and R2

	pc := registers[0]
	// Keep looping, like a physical computer's clock
	for {

		op := memory[pc]
		a1, a2 := memory[pc+1], memory[pc+2]

		// decode and execute
		switch op {
		case Load:
			registers[a1] = memory[a2]
		case Store:
			memory[a2] = registers[a1]
		case Add:
			registers[a1] += registers[a2]
		case Sub:
			registers[a1] -= registers[a2]
		case Addi:
			registers[a1] += a2
		case Subi:
			registers[a1] -= a2
		case Jump:
			pc = a1
			continue
		case Beqz:
			if registers[a1] == 0 {
				pc += a2
			}
		case Halt:
			return
		}

		pc += 3
	}
}
