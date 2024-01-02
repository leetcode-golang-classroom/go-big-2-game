package big_2_game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type SingleShowCardsHandler struct {
	ShowCardHandlerInterface
}

func NewSingleShowCardsHandler(next ShowCardHandlerInterface) ShowCardHandlerInterface {
	singleShowCardHdr := &SingleShowCardsHandler{}
	showCardsHder := NewShowCardsHandler(singleShowCardHdr)
	showCardsHder.SetNext(next)
	return showCardsHder
}

func (singleShowCardHdr *SingleShowCardsHandler) ShowCards(topPlay []*Card, hand []*Card) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	idx := 0
	if len(topPlay) == 0 {
		if strings.Compare(hand[0].String(), "C3") == 0 {
			return "0"
		}
		if len(hand) >= 4 {
			randList := []int{0, 1, len(hand) - 2, len(hand) - 1}
			randIdx := rand.Intn(len(randList))
			idx = randList[randIdx]
		} else {
			idx = rand.Intn(len(hand))
		}
		return fmt.Sprintf("%v", idx)
	}
	lowerBound := singleShowCardHdr.FindLowerbound(topPlay, hand)
	if lowerBound == -1 {
		return "-1"
	}
	totalStep := len(hand) - lowerBound
	if totalStep >= 4 {
		randList := []int{lowerBound, lowerBound + 1, len(hand) - 2, len(hand) - 1}
		randIdx := rand.Intn(len(randList))
		idx = randList[randIdx]
	} else {
		idx = lowerBound + rand.Intn(totalStep)
	}
	return fmt.Sprintf("%v", idx)
}

func (singleShowCardHdr *SingleShowCardsHandler) IsMatch(topPlay []*Card, hand []*Card) bool {
	return (len(topPlay) == 0 || len(topPlay) == 1)
}
func (singleShowCardHdr *SingleShowCardsHandler) FindLowerbound(topPlay []*Card, hand []*Card) int {
	target := topPlay[0]
	for idx, card := range hand {
		if target.IsLess(card) {
			return idx
		}
	}
	return -1
}
