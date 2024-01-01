package big_2_game

import (
	"bufio"
	"strings"
)

type ShowCardFromInputStrategy struct {
	ShowCardStrategy
}

func NewShowCardFromInputStrategy(showCardHdr ShowCardHandlerInterface) ShowCardStrategyInterface {
	return &ShowCardFromInputStrategy{
		ShowCardStrategy: ShowCardStrategy{
			ShowCardHdr: showCardHdr,
		},
	}
}
func (showCardFromInput *ShowCardFromInputStrategy) ShowCards(topPlay []*Card, reader *bufio.Reader, hand []*Card) string {
	lines, _, _ := reader.ReadLine()
	readline := string(lines)
	cardsLine := strings.TrimSpace(readline)
	return cardsLine
}
