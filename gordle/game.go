package gordle

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"io"
	"os"
	"strings"
)

type Game struct {
	solution    []rune
	reader      *bufio.Reader
	maxAttempts int
}

var invalidWordLengthError = fmt.Errorf("invalid guess, word is not the same length as the guess")

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("Roo")
	}
	return nil
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d letter guess: ", len(g.solution))
	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Faild to read input")
			continue
		}
		guess := []rune(string(playerInput))
		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt with gordle is invalid")
		} else {
			return guess
		}
	}
}

func computeFeedback(guess, solution []rune) feedback {
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error, guess and solution lengths differ, %d vs %d", len(guess), len(solution))
		return result
	}

	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			result[posInGuess] = correctPosition
			used[posInGuess] = true
		}
	}

	for posInGuess, character := range guess {
		if result[posInGuess] != absentCharacter {
			continue
		}
		for posInSolution, target := range solution {
			if used[posInSolution] {
				continue
			}
			if character == target {
				result[posInGuess] = wrongPosition
				used[posInSolution] = true
				break
			}
		}
	}
	return result

}

func splitToUpperCaseChars(input string) []rune {
	return []rune(strings.ToUpper(input))
}

func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUpperCaseChars(solution),
		maxAttempts: maxAttempts,
	}
	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")
	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()
		fb := computeFeedback(guess, g.solution)
		fmt.Println(fb)

		if slices.Equal(guess, g.solution) {
			fmt.Printf("You won, you got it in %d guess(s)", currentAttempt)
		}
	}
	fmt.Printf("You lost, the solution was %s \n", string(g.solution))

}
