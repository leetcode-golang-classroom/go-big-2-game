package big_2_game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Deck struct {
	DeckInterface
	cards []*Card
}

type DeckInterface interface {
	Shuffle()
	IsDeckEmpty() bool
	FormatDeck(deckString string)
	Deal(player PlayerInterface)
	DealCard() *Card
	InitDeck()
}

func NewDeck() *Deck {
	return &Deck{
		cards: []*Card{},
	}
}

/*
*
check is deck is Empty
*/
func (deck *Deck) IsDeckEmpty() bool {
	return len(deck.cards) == 0
}

/*
*
use deckString to Initial *Deck
*/
func (deck *Deck) FormatDeck(deckString string) {
	cardStrings := strings.Split(deckString, " ")
	for _, cardString := range cardStrings {
		deck.cards = append(deck.cards, Format(cardString))
	}
}

/*
*
shuffle deck
*/
func (deck *Deck) Shuffle() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(deck.cards), func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	})
}

func (deck *Deck) String() string {
	deckResult := ""
	for idx, card := range deck.cards {
		deckResult += fmt.Sprintf("%v", card)
		if idx != len(deck.cards)-1 {
			deckResult += " "
		}
	}
	return deckResult
}

/*
*
發一張牌到 Player 手牌
*/
func (deck *Deck) Deal(player PlayerInterface) {
	card := deck.DealCard()
	player.AddHand(card)
}

/*
*
從最後一張發牌
*/
func (deck *Deck) DealCard() *Card {
	idx := len(deck.cards) - 1
	card := deck.cards[idx]
	deck.cards = append([]*Card{}, deck.cards[:idx]...)
	return card
}

func (deck *Deck) InitDeck() {
	deck.cards = []*Card{}
	for _, suit := range Suits {
		for _, rank := range Ranks {
			deck.cards = append(deck.cards, NewCard(suit, rank))
		}
	}
}
