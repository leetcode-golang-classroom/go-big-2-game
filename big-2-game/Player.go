package big_2_game

import (
	"fmt"
	"sort"
)

type Player struct {
	PlayerInterface
	name  string
	hands []*Card
}

type PlayerInterface interface {
	NameSelf()
	AddHand(card *Card)
	SortHand()
	IsHandEmpty() bool
	DisplayHand() (string, string)
	GetName() string
	SetName(name string)
	ExtractCards(idxes []int)
	ExtractCard(idx int)
	Play(topPlay []*Card, cardPatternHdr CardPatternHandlerInterface) []*Card
	InitPlay(topPlay []*Card, cardPatternHdr CardPatternHandlerInterface) []*Card
	GetSmallestHand() *Card
}

func NewPlayer() *Player {
	return &Player{
		hands: []*Card{},
		name:  "",
	}
}
func (player *Player) GetName() string {
	return player.name
}
func (player *Player) SetName(name string) {
	player.name = name
}

func (player *Player) AddHand(card *Card) {
	player.hands = append(player.hands, card)
}

func (player *Player) SortHand() {
	sort.Slice(player.hands, func(i, j int) bool {
		return player.hands[i].IsLess(player.hands[j])
	})
}

func (player *Player) IsHandEmpty() bool {
	return len(player.hands) == 0
}

func (player *Player) DisplayHand() (string, string) {
	indexLine := ""
	cardsLine := ""
	for idx, card := range player.hands {
		cardString := fmt.Sprintf("%v", card)
		if idx != len(player.hands)-1 {
			cardString += " "
		}
		indexString := fmt.Sprintf("%v", idx)
		diff := len(cardString) - len(indexString)
		for i := 0; i < diff; i++ {
			indexString += " "
		}
		indexLine += indexString
		cardsLine += cardString
	}
	return indexLine, cardsLine
}

func (player *Player) ExtractCards(idxes []int) {
	newHands := []*Card{}
	idxesMap := make(map[int]struct{})
	for _, idx := range idxes {
		idxesMap[idx] = struct{}{}
	}
	for idx, hand := range player.hands {
		if _, ok := idxesMap[idx]; !ok {
			newHands = append(newHands, hand)
		}
	}
	player.hands = newHands
	player.SortHand()
}

func (player *Player) GetSmallestHand() *Card {
	return player.hands[0]
}
