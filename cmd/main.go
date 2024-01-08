package main

import (
	"bufio"
	"os"

	big_2_game "github.com/leetcode-golang-classroom/go-big-2-game/big-2-game"
)

func main() {
	ioReader := bufio.NewReader(os.Stdin)
	ioWriter := bufio.NewWriter(os.Stdout)
	showCardHdr := big_2_game.NewStraightShowCardsHandler(
		big_2_game.NewFullHouseShowCardsHandler(
			big_2_game.NewPairShowCardsHandler(
				big_2_game.NewPairShowCardsHandler(
					big_2_game.NewSingleShowCardsHandler(
						nil,
					),
				),
			),
		),
	)
	aiShowCardStrategy := big_2_game.NewShowCardFromAIStrategy(showCardHdr)
	showCardStrategy := big_2_game.NewShowCardFromInputStrategy(nil)
	cardPatternHdr := big_2_game.NewSingleCardPatternHandler(
		big_2_game.NewPairCardPatternHandler(
			big_2_game.NewStraightCardPatternHandler(
				big_2_game.NewFullHouseCardPatternHandler(
					nil,
				),
			),
		),
	)

	big2Game := big_2_game.NewGame(cardPatternHdr, []big_2_game.PlayerInterface{
		// big_2_game.NewAIPlayer(ioWriter, aiShowCardStrategy),
		big_2_game.NewHumanPlayer(ioReader, ioWriter, showCardStrategy),
		big_2_game.NewAIPlayer(ioWriter, aiShowCardStrategy),
		big_2_game.NewAIPlayer(ioWriter, aiShowCardStrategy),
		big_2_game.NewAIPlayer(ioWriter, aiShowCardStrategy),
	},
		big_2_game.NewDeck(),
		ioReader,
		big_2_game.NewShuffleFromRandomStrategy(),
		ioWriter,
	)
	big2Game.GameFlow()
}
