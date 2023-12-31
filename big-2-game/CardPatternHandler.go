package big_2_game

import (
	"bufio"
)

type CardPatternHandler struct {
	CardPatternHandlerInterface
	Next CardPatternHandlerInterface
}

type CardPatternHandlerInterface interface {
	SetNext(next CardPatternHandlerInterface)
	IsMatch(topPlay []*Card, show []*Card) bool
	CheckCardPattern(topPlay []*Card, show []*Card, player PlayerInterface, ioWriter *bufio.Writer) bool
	DoCheckCardPattern(topPlay []*Card, show []*Card, player PlayerInterface, ioWriter *bufio.Writer) bool
}

func NewCardPatternHandler(cardPatternHdr CardPatternHandlerInterface) *CardPatternHandler {
	return &CardPatternHandler{
		cardPatternHdr,
		nil,
	}
}
func (cardPatternHdr *CardPatternHandler) SetNext(next CardPatternHandlerInterface) {
	cardPatternHdr.Next = next
}

func (cardPatternHdr *CardPatternHandler) DoCheckCardPattern(topPlay []*Card, show []*Card, player PlayerInterface, ioWriter *bufio.Writer) bool {
	initHdr := cardPatternHdr

	for initHdr != nil {
		if initHdr.IsMatch(topPlay, show) {
			return initHdr.CheckCardPattern(topPlay, show, player, ioWriter)
		} else {
			if initHdr.Next == nil {
				initHdr = nil
			} else {
				initHdr = initHdr.Next.(*CardPatternHandler)
			}
		}
	}
	ioWriter.WriteString("此牌型不合法，請再嘗試一次。\n")
	return false
}
