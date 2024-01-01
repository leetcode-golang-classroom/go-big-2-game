package big_2_game

import "bufio"

type ShowCardFromAIStrategy struct {
	*ShowCardStrategy
}

func NewShowCardFromAIStrategy(showCardHdr ShowCardHandlerInterface) ShowCardStrategyInterface {
	return &ShowCardFromAIStrategy{
		ShowCardStrategy: &ShowCardStrategy{
			ShowCardHdr: showCardHdr,
		},
	}
}

func (showCardFromAI *ShowCardFromAIStrategy) ShowCards(topPlay []*Card, reader *bufio.Reader, hand []*Card) string {
	return showCardFromAI.ShowCardHdr.DoShowCards(topPlay, hand)
}
