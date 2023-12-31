package big_2_game

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type HumanPlayer struct {
	*Player
	ioReader *bufio.Reader
	ioWriter *bufio.Writer
}

func NewHumanPlayer(ioReader *bufio.Reader, ioWriter *bufio.Writer) *HumanPlayer {
	return &HumanPlayer{
		Player:   NewPlayer(),
		ioReader: ioReader,
		ioWriter: ioWriter,
	}
}
func (humanPlayer *HumanPlayer) NameSelf() {
	var name string
	lines, _, _ := humanPlayer.ioReader.ReadLine()
	readline := string(lines)
	name = strings.TrimSpace(readline)
	humanPlayer.SetName(name)
}

func (humanPlayer *HumanPlayer) Play(topPlay []*Card, cardPatternHdr CardPatternHandlerInterface) []*Card {
	isShowCorrect := false
	var readline string
	humanPlayer.DisplayHandAndName()
	result := []*Card{}
	for !isShowCorrect {
		lines, _, _ := humanPlayer.ioReader.ReadLine()
		readline = string(lines)
		cardsLine := strings.TrimSpace(readline)
		isPass := strings.Compare(cardsLine, "-1") == 0
		if isPass && len(topPlay) != 0 {
			humanPlayer.ioWriter.WriteString(fmt.Sprintf("玩家 %v PASS\n", humanPlayer.GetName()))
			isShowCorrect = true
			return []*Card{}
		}
		if isPass && len(topPlay) == 0 {
			humanPlayer.ioWriter.WriteString("你不能在新的回合中喊 PASS\n")
			isShowCorrect = false
			continue
		}
		shows, nIdxes := humanPlayer.ParesInputToShow(cardsLine)
		isShowCorrect = cardPatternHdr.DoCheckCardPattern(topPlay, shows, humanPlayer, humanPlayer.ioWriter)
		if isShowCorrect {
			humanPlayer.ExtractCards(nIdxes)
			result = shows
		}
	}
	return result
}
func (humanPlayer *HumanPlayer) ParesInputToShow(cardsLine string) ([]*Card, []int) {
	idxes := strings.Split(cardsLine, " ")
	shows := []*Card{}
	nIdxes := []int{}
	for _, idx := range idxes {
		num, _ := strconv.Atoi(idx)
		nIdxes = append(nIdxes, num)
		shows = append(shows, humanPlayer.hands[num])
	}
	sort.Slice(shows, func(i, j int) bool {
		return shows[i].IsLess(shows[j])
	})
	return shows, nIdxes
}
func (humanPlayer *HumanPlayer) DisplayHandAndName() {
	humanPlayer.ioWriter.WriteString(fmt.Sprintf("輪到%v了\n", humanPlayer.GetName()))
	indexLine, handsLine := humanPlayer.DisplayHand()
	humanPlayer.ioWriter.WriteString(fmt.Sprintf("%v\n", indexLine))
	humanPlayer.ioWriter.WriteString(fmt.Sprintf("%v\n", handsLine))
}
func (humanPlayer *HumanPlayer) InitPlay(topPlay []*Card, cardPatternHdr CardPatternHandlerInterface) []*Card {
	humanPlayer.DisplayHandAndName()
	isShowCorrect := false
	var readline string
	result := []*Card{}
	for !isShowCorrect {
		lines, _, _ := humanPlayer.ioReader.ReadLine()
		readline = string(lines)
		cardsLine := strings.TrimSpace(readline)
		isPass := strings.Compare(cardsLine, "-1") == 0
		if isPass && len(topPlay) == 0 {
			humanPlayer.ioWriter.WriteString("你不能在新的回合中喊 PASS\n")
			isShowCorrect = false
			continue
		}
		shows, nIdxes := humanPlayer.ParesInputToShow(cardsLine)
		if strings.Compare(shows[0].String(), "C[3]") != 0 {
			humanPlayer.ioWriter.WriteString("此牌型不合法，請再嘗試一次。\n")
			isShowCorrect = false
			continue
		}
		isShowCorrect = cardPatternHdr.DoCheckCardPattern(topPlay, shows, humanPlayer, humanPlayer.ioWriter)
		if isShowCorrect {
			humanPlayer.ExtractCards(nIdxes)
			result = shows
		}
	}
	return result
}
