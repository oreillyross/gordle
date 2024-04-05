package gordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const ErrCorpusIsEmpty = corpusError("corpus is empty")

func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to open file, %q, %w", path, err)
	}
	if len(data) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	words := strings.Fields(string(data))
	fmt.Println(len(words))
	return words, nil
}

func PickWord(corpus []string) string {
	index := rand.Intn(len(corpus))
	return corpus[index]
}
