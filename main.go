package main

import (
	"fmt"
	"learn-pockets/gordle/gordle"
	corpus "learn-pockets/gordle/gordle/corpus"

	"os"
)

func main() {
	wd, werr := os.Getwd()
	fmt.Println(wd, werr)
	corpus, err := corpus.ReadCorpus("gordle/corpus/english.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to read corpus %s", err)
	}

	g := gordle.New(os.Stdin, corpus, 3)
	g.Play()

}
