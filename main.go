package main

import (
	"fmt"

	"github.com/captncraig/zed/zmachine"
)

func main() {

	s, err := zmachine.LoadStory("games/hollywoo.z3")
	fmt.Println(s.Header, err, string(s.Header.Serial[:]))
	fmt.Println(zmachine.Decode(s, uint32(s.Header.StartPC), zmachine.V3Opcodes))
}
