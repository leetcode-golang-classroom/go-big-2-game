package big_2_game

import "bufio"

type ShuffleStrategy interface {
	Shuffle(deck DeckInterface, ioReader *bufio.Reader)
}
