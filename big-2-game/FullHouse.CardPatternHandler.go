package big_2_game

import (
	"bufio"
	"fmt"
	"sort"
)

type FullHouseCardPatternHandler struct {
	CardPatternHandlerInterface
}

func NewFullHouseCardPatternHandler(next CardPatternHandlerInterface) CardPatternHandlerInterface {
	fullHouseCardPatternHdr := &FullHouseCardPatternHandler{}
	cardPatternHdr := NewCardPatternHandler(fullHouseCardPatternHdr)
	cardPatternHdr.SetNext(next)
	return cardPatternHdr
}
func (fullHouseCardPatternHdr *FullHouseCardPatternHandler) IsMatch(topPlay []*Card, show []*Card) bool {
	return (len(topPlay) == 0 || fullHouseCardPatternHdr.IsFullHousePattern(topPlay)) &&
		(fullHouseCardPatternHdr.IsFullHousePattern(show))
}
func (fullHouseCardPatternHdr *FullHouseCardPatternHandler) IsFullHousePattern(cards []*Card) bool {
	rankMap := fullHouseCardPatternHdr.CollectFullHousePattern(cards)
	result := fullHouseCardPatternHdr.FindNumber(rankMap)
	return len(cards) == 5 && len(result) == 2 && result[1] == 3
}
func (fullHouseCardPatternHdr *FullHouseCardPatternHandler) CollectFullHousePattern(cards []*Card) map[int]int {
	rankMap := make(map[int]int)
	for _, card := range cards {
		val, ok := rankMap[int(card.rank)]
		if !ok {
			rankMap[int(card.rank)] = 1
		} else {
			rankMap[int(card.rank)] = val + 1
		}
	}
	return rankMap
}
func (fullHouseCardPatternHdr *FullHouseCardPatternHandler) FindNumber(rankMap map[int]int) []int {
	result := []int{}
	for _, value := range rankMap {
		result = append(result, value)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] <= result[j]
	})
	return result
}

func (fullHouseCardPatternHdr *FullHouseCardPatternHandler) FindFullHouseLargestCard(cards []*Card) *Card {
	rankMap := make(map[Rank]([]*Card))
	result := []*Card{}
	for _, card := range cards {
		val, ok := rankMap[card.rank]
		if !ok {
			rankMap[card.rank] = []*Card{card}
		} else {
			rankMap[card.rank] = append(val, card)
			if len(rankMap[card.rank]) == 3 {
				result = rankMap[card.rank]
				break
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].IsLess(result[j])
	})
	return result[2]
}
func (fullHouseCardPatternHdr *FullHouseCardPatternHandler) CheckCardPattern(topPlay []*Card, show []*Card, player PlayerInterface, ioWriter *bufio.Writer) bool {
	if len(topPlay) == 0 ||
		fullHouseCardPatternHdr.FindFullHouseLargestCard(topPlay).
			IsLess(fullHouseCardPatternHdr.FindFullHouseLargestCard(show)) {
		ioWriter.WriteString(fmt.Sprintf("玩家 %v 打出了 順子 %v %v %v %v %v\n", player.GetName(), show[0], show[1], show[2], show[3], show[4]))
		return true
	}
	ioWriter.WriteString("fullhouse:此牌型不合法，請再嘗試一次。\n")
	return false
}
