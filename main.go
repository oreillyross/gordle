package main

import (
	"learn-pockets/gordle/gordle"
	"os"
)

func main() {
	g := gordle.New(os.Stdin, "roode", 3)
	g.Play()

}
