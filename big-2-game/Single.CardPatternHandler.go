package big_2_game

import (
	"bufio"
	"fmt"
)

type SingleCardPatternHandler struct {
	CardPatternHandlerInterface
}

func NewSingleCardPatternHandler(next CardPatternHandlerInterface) CardPatternHandlerInterface {
	singleCardPatternHdr := &SingleCardPatternHandler{}
	cardPatternHdr := NewCardPatternHandler(singleCardPatternHdr)
	cardPatternHdr.SetNext(next)
	return cardPatternHdr
}
func (singleCardPatternHdr *SingleCardPatternHandler) IsMatch(topPlay []*Card, show []*Card) bool {
	return (len(topPlay) == 0 || len(topPlay) == 1) && (len(show) == 1)
}

func (singleCardPatternHdr *SingleCardPatternHandler) CheckCardPattern(topPlay []*Card, show []*Card, player PlayerInterface, ioWriter *bufio.Writer) bool {
	if len(topPlay) == 0 || topPlay[0].IsLess(show[0]) {
		ioWriter.WriteString(fmt.Sprintf("玩家 %v 打出了 單張 %v\n", player.GetName(), show[0]))
		return true
	}
	ioWriter.WriteString("此牌型不合法，請再嘗試一次。\n")
	return false
}
