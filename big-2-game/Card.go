package big_2_game

import (
	"fmt"
	"strings"
)

type Card struct {
	rank Rank
	suit Suit
}

type Rank int

const (
	Three Rank = iota + 3
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	J
	Q
	K
	A
	Two
)

var ranks = []string{
	"3", "4", "5", "6", "7", "8", "9", "10",
	"J", "Q", "K", "A", "2",
}

type Suit int

const (
	C = iota
	D
	H
	S
)

var suits = []string{
	"C", "D", "H", "S",
}

var Suits = []Suit{
	C,
	D,
	H,
	S,
}
var Ranks = []Rank{
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	J,
	Q,
	K,
	A,
	Two,
}

func (card *Card) String() string {
	return fmt.Sprintf("%v[%v]", suits[card.suit], ranks[card.rank-3])
}

/*
*
check if card is Less than targetCard
*/
func (card *Card) IsLess(targetCard *Card) bool {
	return card.rank < targetCard.rank || (card.rank == targetCard.rank && card.suit < targetCard.suit)
}

/*
*
use cardString to Generate *Card
for example:
cardString = "C[2]" ==> &Card{suit: suit.C, rank: rank.Two}
*
*/
func Format(cardString string) *Card {
	parsedString := strings.Split(cardString, "[")
	suitString := parsedString[0]
	rankString := strings.Split(parsedString[1], "]")[0]
	suitIdx := 0
	rankIdx := 0
	for idx, rank := range ranks {
		if strings.Compare(rank, rankString) == 0 {
			rankIdx = idx
			break
		}
	}
	for idx, suit := range suits {
		if strings.Compare(suit, suitString) == 0 {
			suitIdx = idx
			break
		}
	}
	return &Card{
		rank: Rank(rankIdx + 3),
		suit: Suit(suitIdx),
	}
}

func NewCard(suit Suit, rank Rank) *Card {
	return &Card{
		rank: rank,
		suit: suit,
	}
}
