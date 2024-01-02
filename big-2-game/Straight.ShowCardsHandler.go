package big_2_game

import (
	"fmt"
	"strings"
)

type StraightShowCardsHandler struct {
	ShowCardHandlerInterface
}

func NewStraightShowCardsHandler(next ShowCardHandlerInterface) ShowCardHandlerInterface {
	straightHouseShowCardHdr := &FullHouseShowCardsHandler{}
	showCardsHder := NewShowCardsHandler(straightHouseShowCardHdr)
	showCardsHder.SetNext(next)
	return showCardsHder
}

func (straightShowCardHdr *StraightShowCardsHandler) ShowCards(topPlay []*Card, hand []*Card) string {
	if len(topPlay) == 0 {
		if strings.Compare(hand[0].String(), "C3") == 0 {
			return straightShowCardHdr.FindStraightWithClub3(hand)
		}
		straight := straightShowCardHdr.FindLeastStraigt(hand)
		if len(straight) == 0 {
			return "-1"
		}
		return fmt.Sprintf("%v %v %v %v %v", straight[0], straight[1], straight[2], straight[3], straight[4])
	}
	leadingCardIdx := straightShowCardHdr.FindLeadingCardIdx(topPlay, hand)
	if leadingCardIdx == -1 {
		return "-1"
	}
	followCardsIdx := straightShowCardHdr.FindFollowingCards(leadingCardIdx, hand)
	if len(followCardsIdx) != 4 {
		return "-1"
	}
	return fmt.Sprintf("%v %v %v %v %v", followCardsIdx[0], followCardsIdx[1], followCardsIdx[2], followCardsIdx[3], leadingCardIdx)
}

func (straightShowCardHdr *StraightShowCardsHandler) IsMatch(topPlay []*Card, hand []*Card) bool {
	return (len(topPlay) == 0 || straightShowCardHdr.HasStraightPattern(topPlay)) && straightShowCardHdr.HasStraightPattern(hand)
}
func (straightShowCardHdr *StraightShowCardsHandler) FindStraightWithClub3(hand []*Card) string {
	rankMap := make(map[int][]int)
	for idx := 1; idx < len(hand); idx++ {
		list, ok := rankMap[int(hand[idx].rank)]
		if !ok {
			rankMap[int(hand[idx].rank)] = []int{idx}
		} else {
			rankMap[int(hand[idx].rank)] = append(list, idx)
		}
	}
	result := "0"
	for idx := 1; idx <= 4; idx++ {
		list, ok := rankMap[int(hand[0].rank)+idx]
		if !ok {
			return "-1"
		}
		result += fmt.Sprintf(" %v", list[0])
	}
	return result
}
func (straightShowCardHdr *StraightShowCardsHandler) FindFollowingCards(leadingCardIdx int, hand []*Card) []int {
	rankMap := make(map[int][]int)
	for idx := 0; idx < leadingCardIdx; idx++ {
		list, ok := rankMap[int(hand[idx].rank)]
		if !ok {
			rankMap[int(hand[idx].rank)] = []int{idx}
		} else {
			rankMap[int(hand[idx].rank)] = append(list, idx)
		}
	}
	result := []int{}
	for idx := 1; idx <= 4; idx++ {
		list, ok := rankMap[int(hand[leadingCardIdx].rank)-idx]
		if !ok {
			return []int{}
		}
		result = append(result, list[0])
	}
	return result
}
func (straightShowCardHdr *StraightShowCardsHandler) FindLeadingCardIdx(topPlay []*Card, hand []*Card) int {
	leadingCard := topPlay[4]
	for idx, card := range hand {
		if leadingCard.IsLess(card) {
			return idx
		}
	}
	return -1
}

func (straightShowCardHdr *StraightShowCardsHandler) FindLeastStraigt(hand []*Card) []int {
	if len(hand) < 5 {
		return []int{}
	}
	rankMap := make(map[int][]int)

	for idx, card := range hand {
		list, ok := rankMap[int(card.rank)]
		if !ok {
			rankMap[int(card.rank)] = []int{idx}
		} else {
			rankMap[int(card.rank)] = append(list, idx)
		}
	}
	for rank := range rankMap {
		startRank := rank
		count := 0
		list, ok := rankMap[startRank+count]
		result := []int{}
		for ok {
			count++
			result = append(result, list[0])
			list, ok = rankMap[startRank+count]
			if count == 5 {
				return result
			}
		}
	}
	return []int{}
}
func (straightShowCardHdr *StraightShowCardsHandler) HasStraightPattern(cards []*Card) bool {
	if len(cards) < 5 {
		return false
	}
	rankMap := make(map[int]struct{})

	for _, card := range cards {
		rankMap[int(card.rank)] = struct{}{}
	}
	for rank := range rankMap {
		count := 0
		startRank := rank
		_, ok := rankMap[startRank+count]
		for ok {
			count++
			_, ok = rankMap[startRank+count]
			if count >= 5 {
				return true
			}
		}
	}
	return false
}
