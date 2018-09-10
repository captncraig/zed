package main

import (
	"fmt"

	"github.com/captncraig/zed/zmachine"
)

func main() {
	h, _, err := zmachine.LoadStory("games/hollywoo.z3")
	fmt.Println(h, err, string(h.Serial[:]))
}
