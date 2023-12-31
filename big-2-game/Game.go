package big_2_game

import (
	"bufio"
	"fmt"
	"strings"
)

type Game struct {
	GameInterface
	topPlayer          PlayerInterface
	topPlay            []*Card
	players            []PlayerInterface
	deck               DeckInterface
	round              int
	turn               int
	PassCount          int
	cardPatternHandler CardPatternHandlerInterface
	ioReader           *bufio.Reader
	shuffleStrategy    ShuffleStrategy
	ioWriter           *bufio.Writer
}

type GameInterface interface {
	DealCard(players []PlayerInterface)
	NextTurn()
	NextRound()
	TakeTurn()
	IsTopPlayEmpty() bool
	UpdateTopPlay(cards []*Card)
	IsGameFinished() bool
	PlayGame()
	DisPlayWinner()
	InitTurn()
	prepareGame()
	GameFlow()
	UpdateTopPlayer(player PlayerInterface)
}

func (game *Game) UpdateTopPlayer(player PlayerInterface) {
	game.topPlayer = player
}

func NewGame(cardPatternHandler CardPatternHandlerInterface, players []PlayerInterface, deck DeckInterface, ioReader *bufio.Reader, shuffleStrategy ShuffleStrategy, ioWriter *bufio.Writer) *Game {
	return &Game{
		topPlayer:          nil,
		topPlay:            []*Card{},
		players:            players,
		deck:               deck,
		turn:               0,
		round:              1,
		PassCount:          0,
		cardPatternHandler: cardPatternHandler,
		ioReader:           ioReader,
		shuffleStrategy:    shuffleStrategy,
		ioWriter:           ioWriter,
	}
}

func (game *Game) DealCard(players []PlayerInterface) {
	for !game.deck.IsDeckEmpty() {
		for _, player := range game.players {
			game.deck.Deal(player)
		}
	}
	for _, player := range game.players {
		player.SortHand()
	}
}

func (game *Game) NextTurn() {
	game.turn = ((game.turn + 1) % 4)
}

func (game *Game) NextRound() {
	// clear topPlay
	game.topPlay = []*Card{}
	game.PassCount = 0
	game.round += 1
	game.ioWriter.WriteString("新的回合開始了。\n")
}

func (game *Game) IsTopPlayEmpty() bool {
	return len(game.topPlay) == 0
}

func (game *Game) UpdateTopPlay(cards []*Card) {
	game.topPlay = cards
}

func (game *Game) IsGameFinished() bool {
	return game.topPlayer != nil && game.topPlayer.IsHandEmpty()
}
func (game *Game) PlayGame() {
	game.ioWriter.WriteString("新的回合開始了。\n")
	game.InitTurn()
	for !game.IsGameFinished() {
		game.TakeTurn()
	}
	game.DisPlayWinner()
}
func (game *Game) TakeTurn() {
	player := game.players[game.turn]
	cards := player.Play(game.topPlay, game.cardPatternHandler)
	if len(cards) != 0 {
		game.UpdateTopPlay(cards)
		game.UpdateTopPlayer(player)
		game.PassCount = 0
	} else {
		game.PassCount++
	}
	if player.IsHandEmpty() {
		return
	}
	if game.PassCount == 3 {
		game.NextRound()
	}
	game.NextTurn()
}
func (game *Game) FindInitTopPlayer() PlayerInterface {
	topPlayer := game.players[game.turn]
	for strings.Compare(topPlayer.GetSmallestHand().String(), "C[3]") != 0 {
		game.NextTurn()
		topPlayer = game.players[game.turn]
	}
	return topPlayer
}
func (game *Game) InitTurn() {
	game.UpdateTopPlayer(game.FindInitTopPlayer())
	game.UpdateTopPlay(game.topPlayer.InitPlay(game.topPlay, game.cardPatternHandler))
	game.NextTurn()
}
func (game *Game) prepareGame() {
	game.shuffleStrategy.Shuffle(game.deck, game.ioReader)
	for _, player := range game.players {
		player.NameSelf()
	}
	game.DealCard(game.players)
}
func (game *Game) DisPlayWinner() {
	game.ioWriter.WriteString(fmt.Sprintf("遊戲結束，遊戲的勝利者為 %v\n", game.topPlayer.GetName()))
	game.ioWriter.Flush()
}
func (game *Game) GameFlow() {
	game.prepareGame()
	game.PlayGame()
}
