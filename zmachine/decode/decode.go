package decode

import "github.com/captncraig/zed/zmachine"

type OperandType byte

const (
	OperandTypeLongConstant  OperandType = 0x00
	OperandTypeShortConstant             = 0x01
	OperandTypeVariable                  = 0x02
	OperandTypeOmitted                   = 0x03
)

type Instruction struct {
	OpCount       string // 0OP, 1OP, 2OP or VAR
	OpCode        byte
	OperandTypes  []OperandType
	OperandValues []uint16
	StoreVariable *byte
	BranchOffset  *uint16
	Text          *string

	Raw []byte
}

func Decode(story zmachine.StoryFile, addr uint32) *Instruction {
	instr := &Instruction{}
	b0 := story[addr]
	switch b0 >> 6 {
	case 0x03:
		// variable form
		instr.OpCode = b0 & 0x1F
		if (b0>>5)&1 == 1 {
			instr.OpCount = "VAR"
		} else {
			instr.OpCount = "2OP"
		}
	case 0x02:
		// short form
		instr.OpCode = b0 & 0x0F
		opType := OperandType(b0>>4) & OperandTypeOmitted
		if opType == OperandTypeOmitted {
			instr.OpCount = "0OP"
		} else {
			instr.OpCount = "1OP"
			instr.OperandTypes = []OperandType{opType}
		}
	default:
		// long form EXECPT 0xBE
		instr.OpCode = b0 & 0x1F
		instr.OpCount = "2OP"
		instr.OperandTypes = []OperandType{OperandTypeShortConstant, OperandTypeShortConstant}
		if (b0>>7)&1 == 1 {
			instr.OperandTypes[0] = OperandTypeVariable
		}
		if (b0>>6)&1 == 1 {
			instr.OperandTypes[1] = OperandTypeVariable
		}
	}
}
