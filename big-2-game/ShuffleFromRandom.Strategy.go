package big_2_game

import "bufio"

type ShuffleFromRandomStrategy struct {
	ShuffleStrategy
}

func NewShuffleFromRandomStrategy() ShuffleStrategy {
	return &ShuffleFromRandomStrategy{}
}
func (shuffleRandom *ShuffleFromRandomStrategy) Shuffle(deck DeckInterface, reader *bufio.Reader) {
	deck.InitDeck()
	deck.Shuffle()
}
