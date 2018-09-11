package main

import (
	"fmt"

	"github.com/captncraig/zed/zmachine"
	"github.com/captncraig/zed/zmachine/decode"
)

func main() {
	
	
	
	
	
	
	h, s, err := zmachine.LoadStory("games/hollywoo.z3")
	fmt.Println(h, err, string(h.Serial[:]))
	fmt.Println(decode.Decode(s, uint32(h.StartPC), zmachine.V3Opcodes))
}
