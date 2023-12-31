package test

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	big_2_game "github.com/leetcode-golang-classroom/go-big-2-game/big-2-game"
)

func TestWithTestCase1(t *testing.T) {
	var b bytes.Buffer
	absPath, _ := filepath.Abs("../test_cases/test_case_1.in")
	absOutputPath, _ := filepath.Abs("../test_cases/test_case_1.out")
	file, _ := os.Open(absPath)
	outputFile, _ := os.ReadFile(absOutputPath)
	ioReader := bufio.NewReader(file)
	ioWriter := bufio.NewWriter(&b)
	cardPatternHdr := big_2_game.NewSingleCardPatternHandelr(
		big_2_game.NewPairCardPatternHandler(
			big_2_game.NewStraightCardPatternHandler(
				big_2_game.NewFullHouseCardPatternHandler(
					nil,
				),
			),
		),
	)

	big2Game := big_2_game.NewGame(cardPatternHdr, []big_2_game.PlayerInterface{
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
	},
		big_2_game.NewDeck(),
		ioReader,
		big_2_game.NewShuffleFromFileStragtegy(),
		ioWriter,
	)
	big2Game.GameFlow()
	got := b.String()
	expected := string(outputFile)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
func TestWithFullouse(t *testing.T) {
	var b bytes.Buffer
	absPath, _ := filepath.Abs("../test_cases/fullhouse.in")
	absOutputPath, _ := filepath.Abs("../test_cases/fullhouse.out")
	file, _ := os.Open(absPath)
	outputFile, _ := os.ReadFile(absOutputPath)
	ioReader := bufio.NewReader(file)
	ioWriter := bufio.NewWriter(&b)
	cardPatternHdr := big_2_game.NewSingleCardPatternHandelr(
		big_2_game.NewPairCardPatternHandler(
			big_2_game.NewStraightCardPatternHandler(
				big_2_game.NewFullHouseCardPatternHandler(
					nil,
				),
			),
		),
	)

	big2Game := big_2_game.NewGame(cardPatternHdr, []big_2_game.PlayerInterface{
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
	},
		big_2_game.NewDeck(),
		ioReader,
		big_2_game.NewShuffleFromFileStragtegy(),
		ioWriter,
	)
	big2Game.GameFlow()
	got := b.String()
	expected := string(outputFile)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func TestWithAlwaysPlayFirstCard(t *testing.T) {
	var b bytes.Buffer
	absPath, _ := filepath.Abs("../test_cases/always-play-first-card.in")
	absOutputPath, _ := filepath.Abs("../test_cases/always-play-first-card.out")
	file, _ := os.Open(absPath)
	outputFile, _ := os.ReadFile(absOutputPath)
	ioReader := bufio.NewReader(file)
	ioWriter := bufio.NewWriter(&b)
	cardPatternHdr := big_2_game.NewSingleCardPatternHandelr(
		big_2_game.NewPairCardPatternHandler(
			big_2_game.NewStraightCardPatternHandler(
				big_2_game.NewFullHouseCardPatternHandler(
					nil,
				),
			),
		),
	)

	big2Game := big_2_game.NewGame(cardPatternHdr, []big_2_game.PlayerInterface{
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
	},
		big_2_game.NewDeck(),
		ioReader,
		big_2_game.NewShuffleFromFileStragtegy(),
		ioWriter,
	)
	big2Game.GameFlow()
	got := b.String()
	expected := string(outputFile)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func TestWithIllegalActions(t *testing.T) {
	var b bytes.Buffer
	absPath, _ := filepath.Abs("../test_cases/illegal-actions.in")
	absOutputPath, _ := filepath.Abs("../test_cases/illegal-actions.out")
	file, _ := os.Open(absPath)
	outputFile, _ := os.ReadFile(absOutputPath)
	ioReader := bufio.NewReader(file)
	ioWriter := bufio.NewWriter(&b)
	cardPatternHdr := big_2_game.NewSingleCardPatternHandelr(
		big_2_game.NewPairCardPatternHandler(
			big_2_game.NewStraightCardPatternHandler(
				big_2_game.NewFullHouseCardPatternHandler(
					nil,
				),
			),
		),
	)

	big2Game := big_2_game.NewGame(cardPatternHdr, []big_2_game.PlayerInterface{
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
	},
		big_2_game.NewDeck(),
		ioReader,
		big_2_game.NewShuffleFromFileStragtegy(),
		ioWriter,
	)
	big2Game.GameFlow()
	got := b.String()
	expected := string(outputFile)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func TestWithNormalNoErrorPlay1(t *testing.T) {
	var b bytes.Buffer
	absPath, _ := filepath.Abs("../test_cases/normal-no-error-play1.in")
	absOutputPath, _ := filepath.Abs("../test_cases/normal-no-error-play1.out")
	file, _ := os.Open(absPath)
	outputFile, _ := os.ReadFile(absOutputPath)
	ioReader := bufio.NewReader(file)
	ioWriter := bufio.NewWriter(&b)
	cardPatternHdr := big_2_game.NewSingleCardPatternHandelr(
		big_2_game.NewPairCardPatternHandler(
			big_2_game.NewStraightCardPatternHandler(
				big_2_game.NewFullHouseCardPatternHandler(
					nil,
				),
			),
		),
	)

	big2Game := big_2_game.NewGame(cardPatternHdr, []big_2_game.PlayerInterface{
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
	},
		big_2_game.NewDeck(),
		ioReader,
		big_2_game.NewShuffleFromFileStragtegy(),
		ioWriter,
	)
	big2Game.GameFlow()
	got := b.String()
	expected := string(outputFile)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func TestWithNormalNoErrorPlay2(t *testing.T) {
	var b bytes.Buffer
	absPath, _ := filepath.Abs("../test_cases/normal-no-error-play2.in")
	absOutputPath, _ := filepath.Abs("../test_cases/normal-no-error-play2.out")
	file, _ := os.Open(absPath)
	outputFile, _ := os.ReadFile(absOutputPath)
	ioReader := bufio.NewReader(file)
	ioWriter := bufio.NewWriter(&b)
	cardPatternHdr := big_2_game.NewSingleCardPatternHandelr(
		big_2_game.NewPairCardPatternHandler(
			big_2_game.NewStraightCardPatternHandler(
				big_2_game.NewFullHouseCardPatternHandler(
					nil,
				),
			),
		),
	)

	big2Game := big_2_game.NewGame(cardPatternHdr, []big_2_game.PlayerInterface{
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
		big_2_game.NewHumanPlayer(ioReader, ioWriter),
	},
		big_2_game.NewDeck(),
		ioReader,
		big_2_game.NewShuffleFromFileStragtegy(),
		ioWriter,
	)
	big2Game.GameFlow()
	got := b.String()
	expected := string(outputFile)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
