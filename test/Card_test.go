package test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	big_2_game "github.com/leetcode-golang-classroom/go-big-2-game/big-2-game"
)

func TestCardFormat(t *testing.T) {
	cardStrings := "C[2] S[A]"
	cards := strings.Split(cardStrings, " ")
	cardArray := []*big_2_game.Card{}
	for _, card := range cards {
		cardArray = append(cardArray, big_2_game.Format(card))
	}
	got := fmt.Sprintf("%v", cardArray)
	expected := "[C[2] S[A]]"
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func TestCardIsLess(t *testing.T) {
	cardStrings := "C[2] S[2]"
	cards := strings.Split(cardStrings, " ")
	cardArray := []*big_2_game.Card{}
	for _, card := range cards {
		cardArray = append(cardArray, big_2_game.Format(card))
	}
	got := cardArray[0].IsLess(cardArray[1])
	expected := true
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
