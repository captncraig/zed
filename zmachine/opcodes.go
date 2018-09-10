package zmachine

type Opcode struct {
	Name  string
	Flags OpcodeFlags
}

type OpcodeFlags byte

const (
	FlagStore OpcodeFlags = 1 << iota
	FlagBranch
	FlagText
	None OpcodeFlags = 0
)

var V3Opcodes = map[string]map[byte]Opcode{
	"2OP": map[byte]Opcode{
		0x01: {"je", FlagBranch},
		0x02: {"jl", FlagBranch},
		0x03: {"jg", FlagBranch},
		0x04: {"dec_chk", FlagBranch},
		0x05: {"inc_chk", FlagBranch},
		0x06: {"jin", FlagBranch},
		0x07: {"test", FlagBranch},
		0x08: {"or", FlagStore},
		0x09: {"and", FlagStore},
		0x0A: {"test_attr", FlagBranch},
		0x0B: {"set_attr", None},
		0x0C: {"clear_attr", None},
		0x0D: {"store", None},
		0x0E: {"insert_obj", None},
		0x0F: {"loadw", FlagStore},
		0x10: {"loadb", FlagStore},
		0x11: {"get_prop", FlagStore},
		0x12: {"get_prop_addr", FlagStore},
		0x13: {"get_next_prop", FlagStore},
		0x14: {"add", FlagStore},
		0x15: {"sub", FlagStore},
		0x16: {"mul", FlagStore},
		0x17: {"div", FlagStore},
		0x18: {"mod", FlagStore},
	},
	"1OP": map[byte]Opcode{
		0x00: {"jz", FlagBranch},
		0x01: {"get_sibling", FlagStore | FlagBranch},
		0x02: {"get_child", FlagStore | FlagBranch},
		0x03: {"get_parent", FlagStore},
		0x04: {"get_prop_len", FlagStore},
		0x05: {"inc", None},
		0x06: {"dec", None},
		0x07: {"print_addr", None},
		0x09: {"remove_obj", None},
		0x0A: {"print_obj", None},
		0x0B: {"ret", None},
		0x0C: {"jump", None},
		0x0D: {"print_paddr", None},
		0x0E: {"load", FlagStore},
		0x0F: {"not", FlagStore},
	},
	"0OP": map[byte]Opcode{
		0x00: {"rtrue", None},
		0x01: {"rfalse", None},
		0x02: {"print", None},
		0x03: {"print_ret", None},
		0x04: {"nop", None},
		0x05: {"save", FlagBranch},
		0x06: {"restore", FlagBranch},
		0x07: {"restart", None},
		0x08: {"ret_popper", None},
		0x09: {"pop", None},
		0x0A: {"quit", None},
		0x0B: {"new_line", None},
		0x0C: {"show_status", None},
		0x0D: {"verify", FlagBranch},
	},
	"VAR": map[byte]Opcode{
		0x00: {"call", FlagStore},
		0x01: {"storew", None},
		0x02: {"storeb", None},
		0x03: {"put_prop", None},
		0x04: {"sread", None},
		0x05: {"print_char", None},
		0x06: {"print_num", None},
		0x07: {"random", FlagStore},
		0x08: {"push", None},
		0x09: {"pull", None},
		0x0A: {"split_window", None},
		0x0B: {"set_window", None},
		0x13: {"output_stream", None},
		0x14: {"input_stream", None},
	},
}
