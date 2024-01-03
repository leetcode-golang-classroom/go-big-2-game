package big_2_game

import (
	"fmt"
	"strings"
)

type FullHouseShowCardsHandler struct {
	ShowCardHandlerInterface
}

func NewFullHouseShowCardsHandler(next ShowCardHandlerInterface) ShowCardHandlerInterface {
	fullHouseShowCardHdr := &FullHouseShowCardsHandler{}
	showCardsHder := NewShowCardsHandler(fullHouseShowCardHdr)
	showCardsHder.SetNext(next)
	return showCardsHder
}

func (fullHouseShowCardHdr *FullHouseShowCardsHandler) ShowCards(topPlay []*Card, hand []*Card) string {
	if len(topPlay) == 0 {
		if strings.Compare(hand[0].String(), "C3") == 0 {
			return fullHouseShowCardHdr.FindFullHouseWithClub3(hand)
		}
		return fullHouseShowCardHdr.FindLeastFullHouse(hand)
	}
	leadingCard := fullHouseShowCardHdr.FindFullHouseLeadingCard(topPlay)
	possibleLeadingIdx := fullHouseShowCardHdr.FindLeastIdxOfHand(leadingCard, hand)
	if possibleLeadingIdx == -1 {
		return "-1"
	}
	leadingCards := fullHouseShowCardHdr.FindLeadingCards(possibleLeadingIdx, hand)
	if len(leadingCards) == 0 {
		return "-1"
	}
	followingCards := fullHouseShowCardHdr.FindFollowingCards(leadingCards[0], hand)
	if len(followingCards) == 0 {
		return "-1"
	}
	return fmt.Sprintf("%v %v %v %v %v", followingCards[0], followingCards[1], leadingCards[0], leadingCards[1], leadingCards[2])
}
func (fullHouseShowCardHdr *FullHouseShowCardsHandler) FindFullHouseWithClub3(hand []*Card) string {
	rankMap := make(map[int][]int)
	for idx, card := range hand {
		list, ok := rankMap[int(card.rank)]
		if !ok {
			rankMap[int(card.rank)] = []int{idx}
		} else {
			rankMap[int(card.rank)] = append(list, idx)
		}
	}
	list := rankMap[3]
	if len(list) != 2 && len(list) != 3 {
		return "-1"
	}
	remain := 5 - len(list)
	followingList := []int{}
	for _, list := range rankMap {
		if len(list) == remain {
			followingList = list
		}
	}
	if len(followingList) == 0 {
		return "-1"
	}
	result := ""
	for idx, val := range list {
		result += fmt.Sprintf("%v", val)
		if idx != len(list)-1 {
			result += " "
		}
	}
	for idx, val := range followingList {
		result += fmt.Sprintf("%v", val)
		if idx != len(list)-1 {
			result += " "
		}
	}
	return result
}
func (fullHouseShowCardHdr *FullHouseShowCardsHandler) FindFollowingCards(leadingCardIdx int, hand []*Card) []int {
	rankMap := make(map[Rank][]int)
	leadingCardRank := hand[leadingCardIdx].rank
	for idx, card := range hand {
		if card.rank != leadingCardRank {
			list, ok := rankMap[card.rank]
			if !ok {
				rankMap[card.rank] = []int{idx}
			} else {
				rankMap[card.rank] = append(list, idx)
			}
			list = rankMap[card.rank]
			if len(list) == 2 {
				return list
			}
		}
	}
	return []int{}
}
func (fullHouseShowCardHdr *FullHouseShowCardsHandler) FindLeadingCards(possibleIdx int, hand []*Card) []int {
	rankMap := make(map[Rank][]int)
	for idx := possibleIdx; idx < len(hand); idx++ {
		card := hand[idx]
		list, ok := rankMap[card.rank]
		if !ok {
			rankMap[card.rank] = []int{idx}
		} else {
			rankMap[card.rank] = append(list, idx)
		}
		list = rankMap[card.rank]
		if len(list) == 3 {
			return list
		}
	}
	return []int{}
}
func (fullHouseShowCardHdr *FullHouseShowCardsHandler) IsMatch(topPlay []*Card, hand []*Card) bool {
	return (len(topPlay) == 0 || (fullHouseShowCardHdr.HasFullHousePattern(topPlay))) && fullHouseShowCardHdr.HasFullHousePattern(hand)
}
func (fullHouseShowCardHdr *FullHouseShowCardsHandler) FindLeastIdxOfHand(leadingCard *Card, hand []*Card) int {
	for idx, card := range hand {
		if leadingCard.IsLess(card) {
			return idx
		}
	}
	return -1
}
func (fullHouseShowCardHdr *FullHouseShowCardsHandler) FindFullHouseLeadingCard(cards []*Card) *Card {
	if cards[2].rank == cards[4].rank { // largest on the right
		return cards[4]
	}
	return cards[2]
}
func (fullHouseShowCardHdr *FullHouseShowCardsHandler) FindLeastFullHouse(hand []*Card) string {
	rankMap := make(map[Rank][]int)
	for idx, card := range hand {
		list, ok := rankMap[card.rank]
		if !ok {
			rankMap[card.rank] = []int{idx}
		} else {
			rankMap[card.rank] = append(list, idx)
		}
	}
	valueMap := make(map[int][]int)
	for _, list := range rankMap {
		count := len(list)
		if _, ok := valueMap[count]; !ok {
			valueMap[count] = list
		}
	}
	largeList := valueMap[3]
	smallList := valueMap[2]
	return fmt.Sprintf("%v %v %v %v %v", smallList[0], smallList[1], largeList[0], largeList[1], largeList[2])
}
func (fullHouseShowCardHdr *FullHouseShowCardsHandler) HasFullHousePattern(cards []*Card) bool {
	if len(cards) < 5 {
		return false
	}
	rankMap := make(map[Rank]int)
	for _, card := range cards {
		rankMap[card.rank]++
	}
	valueMap := make(map[int]struct{})
	for _, value := range rankMap {
		if _, ok := valueMap[value]; !ok {
			valueMap[value] = struct{}{}
		}
	}
	_, exist3 := valueMap[3]
	_, exist2 := valueMap[2]
	return exist3 && exist2
}
