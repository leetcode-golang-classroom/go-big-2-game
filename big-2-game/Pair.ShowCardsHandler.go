package big_2_game

import (
	"fmt"
	"strings"
)

type PairShowCardsHandler struct {
	ShowCardHandlerInterface
}

func NewPairShowCardsHandler(next ShowCardHandlerInterface) ShowCardHandlerInterface {
	pairShowCardHdr := &PairShowCardsHandler{}
	showCardsHder := NewShowCardsHandler(pairShowCardHdr)
	showCardsHder.SetNext(next)
	return showCardsHder
}

func (pairShowCardHdr *PairShowCardsHandler) ShowCards(topPlay []*Card, hand []*Card) string {
	if len(topPlay) == 0 {
		if strings.Compare(hand[0].String(), "C3") == 0 {
			return pairShowCardHdr.FoundPairWithClub3(hand)
		}
		return pairShowCardHdr.ReturnLeastPair(hand)
	}
	lowerboundIdx := pairShowCardHdr.FindLeastIdxOfHand(topPlay, hand)
	if lowerboundIdx == -1 {
		return "-1"
	}
	return pairShowCardHdr.FoundMatchAroundLowerBound(lowerboundIdx, hand)
}
func (pairShowCardHdr *PairShowCardsHandler) FoundPairWithClub3(hand []*Card) string {
	if hand[1].rank == hand[0].rank {
		return "0 1"
	}
	return "-1"
}
func (pairShowCardHdr *PairShowCardsHandler) FoundMatchAroundLowerBound(lowerBoundIdx int, hand []*Card) string {
	rankMap := make(map[Rank][]int)
	for idx := lowerBoundIdx; idx < len(hand); idx++ {
		list, ok := rankMap[hand[idx].rank]
		if !ok {
			rankMap[hand[idx].rank] = []int{idx}
		} else {
			rankMap[hand[idx].rank] = append(list, idx)
		}
		list = rankMap[hand[idx].rank]
		if len(list) == 2 {
			return fmt.Sprintf("%v %v", list[0], list[1])
		}
	}
	return "-1"
}
func (pairShowCardHdr *PairShowCardsHandler) IsMatch(topPlay []*Card, hand []*Card) bool {
	return (len(topPlay) == 0 || (len(topPlay) == 2 && topPlay[0].rank == topPlay[1].rank)) && (len(hand) >= 2 && pairShowCardHdr.HasPairPattern(hand))
}
func (pairShowCardHdr *PairShowCardsHandler) FindLeastIdxOfHand(topPlay []*Card, hand []*Card) int {
	target := topPlay[1]
	for idx, card := range hand {
		if target.IsLess(card) {
			return idx
		}
	}
	return -1
}
func (pairShowCardHdr *PairShowCardsHandler) ReturnLeastPair(hand []*Card) string {
	rankMap := make(map[Rank]int)
	for idx, card := range hand {
		foundIdx, ok := rankMap[card.rank]
		if ok {
			return fmt.Sprintf("%v %v", foundIdx, idx)
		} else {
			rankMap[card.rank] = idx
		}
	}
	return "-1"
}
func (pairShowCardHdr *PairShowCardsHandler) HasPairPattern(hand []*Card) bool {
	rankMap := make(map[Rank]int)
	maxCount := 0
	for _, card := range hand {
		count, ok := rankMap[card.rank]
		if ok {
			return true
		}
		count++
		if maxCount < count {
			maxCount = count
		}
		rankMap[card.rank] = count
	}
	return maxCount >= 2
}
