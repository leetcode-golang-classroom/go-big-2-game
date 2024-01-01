package big_2_game

import "bufio"

type ShowCardStrategyInterface interface {
	ShowCards(topPlay []*Card, reader *bufio.Reader, hand []*Card) string
}

type ShowCardStrategy struct {
	ShowCardStrategyInterface
	ShowCardHdr ShowCardHandlerInterface
}
