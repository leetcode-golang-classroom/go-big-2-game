package big_2_game

import (
	"bufio"
	"strings"
)

type ShuffleFromFileStrategy struct {
	ShuffleStrategy
}

func NewShuffleFromFileStrategy() ShuffleStrategy {
	return &ShuffleFromFileStrategy{}
}
func (shufflefile *ShuffleFromFileStrategy) Shuffle(deck DeckInterface, reader *bufio.Reader) {
	readline, _, _ := reader.ReadLine()
	deckString := string(readline)
	deckString = strings.TrimSpace(deckString)
	deck.FormatDeck(deckString)
}
