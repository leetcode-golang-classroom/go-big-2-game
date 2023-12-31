package big_2_game

import (
	"bufio"
	"fmt"
)

type StraightCardPatternHandler struct {
	CardPatternHandlerInterface
}

func NewStraightCardPatternHandler(next CardPatternHandlerInterface) CardPatternHandlerInterface {
	straightCardPatternHdr := &StraightCardPatternHandler{}
	cardPatternHdr := NewCardPatternHandler(straightCardPatternHdr)
	cardPatternHdr.SetNext(next)
	return cardPatternHdr
}
func (straightCardPatternHdr *StraightCardPatternHandler) IsMatch(topPlay []*Card, show []*Card) bool {
	return (len(topPlay) == 0 || (len(topPlay) == 5 &&
		topPlay[0].rank+1 == topPlay[1].rank &&
		topPlay[1].rank+1 == topPlay[2].rank &&
		topPlay[2].rank+1 == topPlay[3].rank &&
		topPlay[3].rank+1 == topPlay[4].rank)) &&
		(len(show) == 5 &&
			show[0].rank+1 == show[1].rank &&
			show[1].rank+1 == show[2].rank &&
			show[2].rank+1 == show[3].rank &&
			show[3].rank+1 == show[4].rank)
}

func (straightCardPatternHdr *StraightCardPatternHandler) CheckCardPattern(topPlay []*Card, show []*Card, player PlayerInterface, ioWriter *bufio.Writer) bool {
	if len(topPlay) == 0 || topPlay[4].IsLess(show[4]) {
		ioWriter.WriteString(fmt.Sprintf("玩家 %v 打出了 順子 %v %v %v %v %v\n", player.GetName(), show[0], show[1], show[2], show[3], show[4]))
		return true
	}
	ioWriter.WriteString("straight:此牌型不合法，請再嘗試一次。\n")
	return false
}
