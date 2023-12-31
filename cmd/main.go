package main

import (
	"bufio"
	"os"

	big_2_game "github.com/leetcode-golang-classroom/go-big-2-game/big-2-game"
)

func main() {
	ioReader := bufio.NewReader(os.Stdin)
	ioWriter := bufio.NewWriter(os.Stdout)
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
}
