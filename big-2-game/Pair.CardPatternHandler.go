package big_2_game

import (
	"bufio"
	"fmt"
)

type PairCardPatternHandler struct {
	CardPatternHandlerInterface
}

func NewPairCardPatternHandler(next CardPatternHandlerInterface) CardPatternHandlerInterface {
	pairCardPatternHdr := &PairCardPatternHandler{}
	cardPatternHdr := NewCardPatternHandler(pairCardPatternHdr)
	cardPatternHdr.SetNext(next)
	return cardPatternHdr
}
func (pairCardPatternHdr *PairCardPatternHandler) IsMatch(topPlay []*Card, show []*Card) bool {
	return (len(topPlay) == 0 || (len(topPlay) == 2 && topPlay[0].rank == topPlay[1].rank)) &&
		(len(show) == 2 && show[0].rank == show[1].rank)
}

func (pairCardPatternHdr *PairCardPatternHandler) CheckCardPattern(topPlay []*Card, show []*Card, player PlayerInterface, ioWriter *bufio.Writer) bool {
	if len(topPlay) == 0 || topPlay[1].IsLess(show[1]) {
		ioWriter.WriteString(fmt.Sprintf("玩家 %v 打出了 對子 %v %v\n", player.GetName(), show[0], show[1]))
		return true
	}
	ioWriter.WriteString("此牌型不合法，請再嘗試一次。\n")
	return false
}
