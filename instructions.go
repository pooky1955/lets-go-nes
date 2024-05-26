package main

type Instruction struct {
	Opcode    int
	Address   string
	Cycles    int
	Operation string
}

const (
	OP_ADC  = "ADC"
	OP_AND  = "AND"
	OP_ASL  = "ASL"
	OP_BCC  = "BCC"
	OP_BCS  = "BCS"
	OP_BEQ  = "BEQ"
	OP_BIT  = "BIT"
	OP_BMI  = "BMI"
	OP_BNE  = "BNE"
	OP_BPL  = "BPL"
	OP_BRK  = "BRK"
	OP_BVC  = "BVC"
	OP_BVS  = "BVS"
	OP_CLC  = "CLC"
	OP_CLD  = "CLD"
	OP_CLI  = "CLI"
	OP_CLV  = "CLV"
	OP_CMP  = "CMP"
	OP_CPX  = "CPX"
	OP_CPY  = "CPY"
	OP_DEC  = "DEC"
	OP_DEX  = "DEX"
	OP_DEY  = "DEY"
	OP_EOR  = "EOR"
	OP_INC  = "INC"
	OP_INX  = "INX"
	OP_INY  = "INY"
	OP_JMP  = "JMP"
	OP_JSR  = "JSR"
	OP_LDA  = "LDA"
	OP_LDX  = "LDX"
	OP_LDY  = "LDY"
	OP_LSR  = "LSR"
	OP_NOP  = "NOP"
	OP_ORA  = "ORA"
	OP_PHA  = "PHA"
	OP_PHP  = "PHP"
	OP_PLA  = "PLA"
	OP_PLP  = "PLP"
	OP_ROL  = "ROL"
	OP_ROR  = "ROR"
	OP_RTI  = "RTI"
	OP_RTS  = "RTS"
	OP_SBC  = "SBC"
	OP_SEC  = "SEC"
	OP_SED  = "SED"
	OP_SEI  = "SEI"
	OP_STA  = "STA"
	OP_STX  = "STX"
	OP_STY  = "STY"
	OP_TAX  = "TAX"
	OP_TAY  = "TAY"
	OP_TSX  = "TSX"
	OP_TXA  = "TXA"
	OP_TXS  = "TXS"
	OP_TYA  = "TYA"
	ADR_ABS = "ABS"
	ADR_ABX = "ABX"
	ADR_ABY = "ABY"
	ADR_ACC = "ACC"
	ADR_IDX = "IDX"
	ADR_IDY = "IDY"
	ADR_IMM = "IMM"
	ADR_IMP = "IMP"
	ADR_IND = "IND"
	ADR_REL = "REL"
	ADR_ZP  = "ZP"
	ADR_ZPX = "ZPX"
	ADR_ZPY = "ZPY"
)

var instructionMap map[int]Instruction = map[int]Instruction{
	0x69: Instruction{Opcode: 0x69, Address: ADR_IMM, Operation: OP_ADC, Cycles: 2},
	0x65: Instruction{Opcode: 0x65, Address: ADR_ZP, Operation: OP_ADC, Cycles: 3},
	0x75: Instruction{Opcode: 0x75, Address: ADR_ZPX, Operation: OP_ADC, Cycles: 4},
	0x6D: Instruction{Opcode: 0x6D, Address: ADR_ABS, Operation: OP_ADC, Cycles: 4},
	0x7D: Instruction{Opcode: 0x7D, Address: ADR_ABX, Operation: OP_ADC, Cycles: 4},
	0x79: Instruction{Opcode: 0x79, Address: ADR_ABY, Operation: OP_ADC, Cycles: 4},
	0x61: Instruction{Opcode: 0x61, Address: ADR_IDX, Operation: OP_ADC, Cycles: 6},
	0x71: Instruction{Opcode: 0x71, Address: ADR_IDY, Operation: OP_ADC, Cycles: 5},
	0x29: Instruction{Opcode: 0x29, Address: ADR_IMM, Operation: OP_AND, Cycles: 2},
	0x25: Instruction{Opcode: 0x25, Address: ADR_ZP, Operation: OP_AND, Cycles: 3},
	0x35: Instruction{Opcode: 0x35, Address: ADR_ZPX, Operation: OP_AND, Cycles: 4},
	0x2D: Instruction{Opcode: 0x2D, Address: ADR_ABS, Operation: OP_AND, Cycles: 4},
	0x3D: Instruction{Opcode: 0x3D, Address: ADR_ABX, Operation: OP_AND, Cycles: 4},
	0x39: Instruction{Opcode: 0x39, Address: ADR_ABY, Operation: OP_AND, Cycles: 4},
	0x21: Instruction{Opcode: 0x21, Address: ADR_IDX, Operation: OP_AND, Cycles: 6},
	0x31: Instruction{Opcode: 0x31, Address: ADR_IDY, Operation: OP_AND, Cycles: 5},
	0x0A: Instruction{Opcode: 0x0A, Address: ADR_ACC, Operation: OP_ASL, Cycles: 2},
	0x06: Instruction{Opcode: 0x06, Address: ADR_ZP, Operation: OP_ASL, Cycles: 5},
	0x16: Instruction{Opcode: 0x16, Address: ADR_ZPX, Operation: OP_ASL, Cycles: 6},
	0x0E: Instruction{Opcode: 0x0E, Address: ADR_ABS, Operation: OP_ASL, Cycles: 6},
	0x1E: Instruction{Opcode: 0x1E, Address: ADR_ABX, Operation: OP_ASL, Cycles: 7},
	0x90: Instruction{Opcode: 0x90, Address: ADR_REL, Operation: OP_BCC, Cycles: 2},
	0xB0: Instruction{Opcode: 0xB0, Address: ADR_REL, Operation: OP_BCS, Cycles: 2},
	0xF0: Instruction{Opcode: 0xF0, Address: ADR_REL, Operation: OP_BEQ, Cycles: 2},
	0x24: Instruction{Opcode: 0x24, Address: ADR_ZP, Operation: OP_BIT, Cycles: 3},
	0x2C: Instruction{Opcode: 0x2C, Address: ADR_ABS, Operation: OP_BIT, Cycles: 4},
	0x30: Instruction{Opcode: 0x30, Address: ADR_REL, Operation: OP_BMI, Cycles: 2},
	0xD0: Instruction{Opcode: 0xD0, Address: ADR_REL, Operation: OP_BNE, Cycles: 2},
	0x10: Instruction{Opcode: 0x10, Address: ADR_REL, Operation: OP_BPL, Cycles: 2},
	0x00: Instruction{Opcode: 0x00, Address: ADR_IMP, Operation: OP_BRK, Cycles: 7},
	0x50: Instruction{Opcode: 0x50, Address: ADR_REL, Operation: OP_BVC, Cycles: 2},
	0x70: Instruction{Opcode: 0x70, Address: ADR_REL, Operation: OP_BVS, Cycles: 2},
	0x18: Instruction{Opcode: 0x18, Address: ADR_IMP, Operation: OP_CLC, Cycles: 2},
	0xD8: Instruction{Opcode: 0xD8, Address: ADR_IMP, Operation: OP_CLD, Cycles: 2},
	0x58: Instruction{Opcode: 0x58, Address: ADR_IMP, Operation: OP_CLI, Cycles: 2},
	0xB8: Instruction{Opcode: 0xB8, Address: ADR_IMP, Operation: OP_CLV, Cycles: 2},
	0xC9: Instruction{Opcode: 0xC9, Address: ADR_IMM, Operation: OP_CMP, Cycles: 2},
	0xC5: Instruction{Opcode: 0xC5, Address: ADR_ZP, Operation: OP_CMP, Cycles: 3},
	0xD5: Instruction{Opcode: 0xD5, Address: ADR_ZPX, Operation: OP_CMP, Cycles: 4},
	0xCD: Instruction{Opcode: 0xCD, Address: ADR_ABS, Operation: OP_CMP, Cycles: 4},
	0xDD: Instruction{Opcode: 0xDD, Address: ADR_ABX, Operation: OP_CMP, Cycles: 4},
	0xD9: Instruction{Opcode: 0xD9, Address: ADR_ABY, Operation: OP_CMP, Cycles: 4},
	0xC1: Instruction{Opcode: 0xC1, Address: ADR_IDX, Operation: OP_CMP, Cycles: 6},
	0xD1: Instruction{Opcode: 0xD1, Address: ADR_IDY, Operation: OP_CMP, Cycles: 5},
	0xE0: Instruction{Opcode: 0xE0, Address: ADR_IMM, Operation: OP_CPX, Cycles: 2},
	0xE4: Instruction{Opcode: 0xE4, Address: ADR_ZP, Operation: OP_CPX, Cycles: 3},
	0xEC: Instruction{Opcode: 0xEC, Address: ADR_ABS, Operation: OP_CPX, Cycles: 4},
	0xC0: Instruction{Opcode: 0xC0, Address: ADR_IMM, Operation: OP_CPY, Cycles: 2},
	0xC4: Instruction{Opcode: 0xC4, Address: ADR_ZP, Operation: OP_CPY, Cycles: 3},
	0xCC: Instruction{Opcode: 0xCC, Address: ADR_ABS, Operation: OP_CPY, Cycles: 4},
	0xC6: Instruction{Opcode: 0xC6, Address: ADR_ZP, Operation: OP_DEC, Cycles: 5},
	0xD6: Instruction{Opcode: 0xD6, Address: ADR_ZPX, Operation: OP_DEC, Cycles: 6},
	0xCE: Instruction{Opcode: 0xCE, Address: ADR_ABS, Operation: OP_DEC, Cycles: 6},
	0xDE: Instruction{Opcode: 0xDE, Address: ADR_ABX, Operation: OP_DEC, Cycles: 7},
	0xCA: Instruction{Opcode: 0xCA, Address: ADR_IMP, Operation: OP_DEX, Cycles: 2},
	0x88: Instruction{Opcode: 0x88, Address: ADR_IMP, Operation: OP_DEY, Cycles: 2},
	0x49: Instruction{Opcode: 0x49, Address: ADR_IMM, Operation: OP_EOR, Cycles: 2},
	0x45: Instruction{Opcode: 0x45, Address: ADR_ZP, Operation: OP_EOR, Cycles: 3},
	0x55: Instruction{Opcode: 0x55, Address: ADR_ZPX, Operation: OP_EOR, Cycles: 4},
	0x4D: Instruction{Opcode: 0x4D, Address: ADR_ABS, Operation: OP_EOR, Cycles: 4},
	0x5D: Instruction{Opcode: 0x5D, Address: ADR_ABX, Operation: OP_EOR, Cycles: 4},
	0x59: Instruction{Opcode: 0x59, Address: ADR_ABY, Operation: OP_EOR, Cycles: 4},
	0x41: Instruction{Opcode: 0x41, Address: ADR_IDX, Operation: OP_EOR, Cycles: 6},
	0x51: Instruction{Opcode: 0x51, Address: ADR_IDY, Operation: OP_EOR, Cycles: 5},
	0xE6: Instruction{Opcode: 0xE6, Address: ADR_ZP, Operation: OP_INC, Cycles: 5},
	0xF6: Instruction{Opcode: 0xF6, Address: ADR_ZPX, Operation: OP_INC, Cycles: 6},
	0xEE: Instruction{Opcode: 0xEE, Address: ADR_ABS, Operation: OP_INC, Cycles: 6},
	0xFE: Instruction{Opcode: 0xFE, Address: ADR_ABX, Operation: OP_INC, Cycles: 7},
	0xE8: Instruction{Opcode: 0xE8, Address: ADR_IMP, Operation: OP_INX, Cycles: 2},
	0xC8: Instruction{Opcode: 0xC8, Address: ADR_IMP, Operation: OP_INY, Cycles: 2},
	0x4C: Instruction{Opcode: 0x4C, Address: ADR_ABS, Operation: OP_JMP, Cycles: 3},
	0x6C: Instruction{Opcode: 0x6C, Address: ADR_IND, Operation: OP_JMP, Cycles: 5},
	0x20: Instruction{Opcode: 0x20, Address: ADR_ABS, Operation: OP_JSR, Cycles: 6},
	0xA9: Instruction{Opcode: 0xA9, Address: ADR_IMM, Operation: OP_LDA, Cycles: 2},
	0xA5: Instruction{Opcode: 0xA5, Address: ADR_ZP, Operation: OP_LDA, Cycles: 3},
	0xB5: Instruction{Opcode: 0xB5, Address: ADR_ZPX, Operation: OP_LDA, Cycles: 4},
	0xAD: Instruction{Opcode: 0xAD, Address: ADR_ABS, Operation: OP_LDA, Cycles: 4},
	0xBD: Instruction{Opcode: 0xBD, Address: ADR_ABX, Operation: OP_LDA, Cycles: 4},
	0xB9: Instruction{Opcode: 0xB9, Address: ADR_ABY, Operation: OP_LDA, Cycles: 4},
	0xA1: Instruction{Opcode: 0xA1, Address: ADR_IDX, Operation: OP_LDA, Cycles: 6},
	0xB1: Instruction{Opcode: 0xB1, Address: ADR_IDY, Operation: OP_LDA, Cycles: 5},
	0xA2: Instruction{Opcode: 0xA2, Address: ADR_IMM, Operation: OP_LDX, Cycles: 2},
	0xA6: Instruction{Opcode: 0xA6, Address: ADR_ZP, Operation: OP_LDX, Cycles: 3},
	0xB6: Instruction{Opcode: 0xB6, Address: ADR_ZPY, Operation: OP_LDX, Cycles: 4},
	0xAE: Instruction{Opcode: 0xAE, Address: ADR_ABS, Operation: OP_LDX, Cycles: 4},
	0xBE: Instruction{Opcode: 0xBE, Address: ADR_ABY, Operation: OP_LDX, Cycles: 4},
	0xA0: Instruction{Opcode: 0xA0, Address: ADR_IMM, Operation: OP_LDY, Cycles: 2},
	0xA4: Instruction{Opcode: 0xA4, Address: ADR_ZP, Operation: OP_LDY, Cycles: 3},
	0xB4: Instruction{Opcode: 0xB4, Address: ADR_ZPX, Operation: OP_LDY, Cycles: 4},
	0xAC: Instruction{Opcode: 0xAC, Address: ADR_ABS, Operation: OP_LDY, Cycles: 4},
	0xBC: Instruction{Opcode: 0xBC, Address: ADR_ABX, Operation: OP_LDY, Cycles: 4},
	0x4A: Instruction{Opcode: 0x4A, Address: ADR_ACC, Operation: OP_LSR, Cycles: 2},
	0x46: Instruction{Opcode: 0x46, Address: ADR_ZP, Operation: OP_LSR, Cycles: 5},
	0x56: Instruction{Opcode: 0x56, Address: ADR_ZPX, Operation: OP_LSR, Cycles: 6},
	0x4E: Instruction{Opcode: 0x4E, Address: ADR_ABS, Operation: OP_LSR, Cycles: 6},
	0x5E: Instruction{Opcode: 0x5E, Address: ADR_ABX, Operation: OP_LSR, Cycles: 7},
	0xEA: Instruction{Opcode: 0xEA, Address: ADR_IMP, Operation: OP_NOP, Cycles: 2},
	0x09: Instruction{Opcode: 0x09, Address: ADR_IMM, Operation: OP_ORA, Cycles: 2},
	0x05: Instruction{Opcode: 0x05, Address: ADR_ZP, Operation: OP_ORA, Cycles: 3},
	0x15: Instruction{Opcode: 0x15, Address: ADR_ZPX, Operation: OP_ORA, Cycles: 4},
	0x0D: Instruction{Opcode: 0x0D, Address: ADR_ABS, Operation: OP_ORA, Cycles: 4},
	0x1D: Instruction{Opcode: 0x1D, Address: ADR_ABX, Operation: OP_ORA, Cycles: 4},
	0x19: Instruction{Opcode: 0x19, Address: ADR_ABY, Operation: OP_ORA, Cycles: 4},
	0x01: Instruction{Opcode: 0x01, Address: ADR_IDX, Operation: OP_ORA, Cycles: 6},
	0x11: Instruction{Opcode: 0x11, Address: ADR_IDY, Operation: OP_ORA, Cycles: 5},
	0x48: Instruction{Opcode: 0x48, Address: ADR_IMP, Operation: OP_PHA, Cycles: 3},
	0x08: Instruction{Opcode: 0x08, Address: ADR_IMP, Operation: OP_PHP, Cycles: 3},
	0x68: Instruction{Opcode: 0x68, Address: ADR_IMP, Operation: OP_PLA, Cycles: 4},
	0x28: Instruction{Opcode: 0x28, Address: ADR_IMP, Operation: OP_PLP, Cycles: 4},
	0x2A: Instruction{Opcode: 0x2A, Address: ADR_ACC, Operation: OP_ROL, Cycles: 2},
	0x26: Instruction{Opcode: 0x26, Address: ADR_ZP, Operation: OP_ROL, Cycles: 5},
	0x36: Instruction{Opcode: 0x36, Address: ADR_ZPX, Operation: OP_ROL, Cycles: 6},
	0x2E: Instruction{Opcode: 0x2E, Address: ADR_ABS, Operation: OP_ROL, Cycles: 6},
	0x3E: Instruction{Opcode: 0x3E, Address: ADR_ABX, Operation: OP_ROL, Cycles: 7},
	0x6A: Instruction{Opcode: 0x6A, Address: ADR_ACC, Operation: OP_ROR, Cycles: 2},
	0x66: Instruction{Opcode: 0x66, Address: ADR_ZP, Operation: OP_ROR, Cycles: 5},
	0x76: Instruction{Opcode: 0x76, Address: ADR_ZPX, Operation: OP_ROR, Cycles: 6},
	0x6E: Instruction{Opcode: 0x6E, Address: ADR_ABS, Operation: OP_ROR, Cycles: 6},
	0x7E: Instruction{Opcode: 0x7E, Address: ADR_ABX, Operation: OP_ROR, Cycles: 7},
	0x40: Instruction{Opcode: 0x40, Address: ADR_IMP, Operation: OP_RTI, Cycles: 6},
	0x60: Instruction{Opcode: 0x60, Address: ADR_IMP, Operation: OP_RTS, Cycles: 6},
	0xE9: Instruction{Opcode: 0xE9, Address: ADR_IMM, Operation: OP_SBC, Cycles: 2},
	0xE5: Instruction{Opcode: 0xE5, Address: ADR_ZP, Operation: OP_SBC, Cycles: 3},
	0xF5: Instruction{Opcode: 0xF5, Address: ADR_ZPX, Operation: OP_SBC, Cycles: 4},
	0xED: Instruction{Opcode: 0xED, Address: ADR_ABS, Operation: OP_SBC, Cycles: 4},
	0xFD: Instruction{Opcode: 0xFD, Address: ADR_ABX, Operation: OP_SBC, Cycles: 4},
	0xF9: Instruction{Opcode: 0xF9, Address: ADR_ABY, Operation: OP_SBC, Cycles: 4},
	0xE1: Instruction{Opcode: 0xE1, Address: ADR_IDX, Operation: OP_SBC, Cycles: 6},
	0xF1: Instruction{Opcode: 0xF1, Address: ADR_IDY, Operation: OP_SBC, Cycles: 5},
	0x38: Instruction{Opcode: 0x38, Address: ADR_IMP, Operation: OP_SEC, Cycles: 2},
	0xF8: Instruction{Opcode: 0xF8, Address: ADR_IMP, Operation: OP_SED, Cycles: 2},
	0x78: Instruction{Opcode: 0x78, Address: ADR_IMP, Operation: OP_SEI, Cycles: 2},
	0x85: Instruction{Opcode: 0x85, Address: ADR_ZP, Operation: OP_STA, Cycles: 3},
	0x95: Instruction{Opcode: 0x95, Address: ADR_ZPX, Operation: OP_STA, Cycles: 4},
	0x8D: Instruction{Opcode: 0x8D, Address: ADR_ABS, Operation: OP_STA, Cycles: 4},
	0x9D: Instruction{Opcode: 0x9D, Address: ADR_ABX, Operation: OP_STA, Cycles: 5},
	0x99: Instruction{Opcode: 0x99, Address: ADR_ABY, Operation: OP_STA, Cycles: 5},
	0x81: Instruction{Opcode: 0x81, Address: ADR_IDX, Operation: OP_STA, Cycles: 6},
	0x91: Instruction{Opcode: 0x91, Address: ADR_IDY, Operation: OP_STA, Cycles: 6},
	0x86: Instruction{Opcode: 0x86, Address: ADR_ZP, Operation: OP_STX, Cycles: 3},
	0x96: Instruction{Opcode: 0x96, Address: ADR_ZPY, Operation: OP_STX, Cycles: 4},
	0x8E: Instruction{Opcode: 0x8E, Address: ADR_ABS, Operation: OP_STX, Cycles: 4},
	0x84: Instruction{Opcode: 0x84, Address: ADR_ZP, Operation: OP_STY, Cycles: 3},
	0x94: Instruction{Opcode: 0x94, Address: ADR_ZPX, Operation: OP_STY, Cycles: 4},
	0x8C: Instruction{Opcode: 0x8C, Address: ADR_ABS, Operation: OP_STY, Cycles: 4},
	0xAA: Instruction{Opcode: 0xAA, Address: ADR_IMP, Operation: OP_TAX, Cycles: 2},
	0xA8: Instruction{Opcode: 0xA8, Address: ADR_IMP, Operation: OP_TAY, Cycles: 2},
	0xBA: Instruction{Opcode: 0xBA, Address: ADR_IMP, Operation: OP_TSX, Cycles: 2},
	0x8A: Instruction{Opcode: 0x8A, Address: ADR_IMP, Operation: OP_TXA, Cycles: 2},
	0x9A: Instruction{Opcode: 0x9A, Address: ADR_IMP, Operation: OP_TXS, Cycles: 2},
	0x98: Instruction{Opcode: 0x98, Address: ADR_IMP, Operation: OP_TYA, Cycles: 2},
}
