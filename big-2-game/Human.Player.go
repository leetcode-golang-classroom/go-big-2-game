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

func NewHumanPlayer(ioReader *bufio.Reader, ioWriter *bufio.Writer, showCardStrategy ShowCardStrategyInterface) *HumanPlayer {
	return &HumanPlayer{
		Player:   NewPlayer(showCardStrategy),
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
	humanPlayer.DisplayName()
	result := []*Card{}
	for !isShowCorrect {
		humanPlayer.DisplayOnlyHand()
		humanPlayer.ioWriter.Flush()
		cardsLine := humanPlayer.ShowCardStrategy.ShowCards(topPlay, humanPlayer.ioReader, humanPlayer.hands)
		isPass := strings.Compare(cardsLine, "-1") == 0
		if isPass && len(topPlay) != 0 {
			humanPlayer.ioWriter.WriteString(fmt.Sprintf("玩家 %v PASS.\n", humanPlayer.GetName()))
			humanPlayer.ioWriter.Flush()
			isShowCorrect = true
			return []*Card{}
		}
		if isPass && len(topPlay) == 0 {
			humanPlayer.ioWriter.WriteString("你不能在新的回合中喊 PASS\n")
			humanPlayer.ioWriter.Flush()
			isShowCorrect = false
			continue
		}
		shows, nIdxes, err := humanPlayer.ParesInputToShow(cardsLine)
		if err != nil {
			isShowCorrect = false
			continue
		}
		isShowCorrect = cardPatternHdr.DoCheckCardPattern(topPlay, shows, humanPlayer, humanPlayer.ioWriter)
		if isShowCorrect {
			humanPlayer.ExtractCards(nIdxes)
			result = shows
		}
		humanPlayer.ioWriter.Flush()
	}
	return result
}
func (humanPlayer *HumanPlayer) ParesInputToShow(cardsLine string) ([]*Card, []int, error) {
	idxes := strings.Split(cardsLine, " ")
	shows := []*Card{}
	nIdxes := []int{}
	for _, idx := range idxes {
		num, err := strconv.Atoi(idx)
		if err != nil {
			return nil, nil, err
		}
		if num < -1 || num > len(humanPlayer.hands)-1 {
			return nil, nil, fmt.Errorf("%v out of index range", num)
		}
		nIdxes = append(nIdxes, num)
		shows = append(shows, humanPlayer.hands[num])
	}
	sort.Slice(shows, func(i, j int) bool {
		return shows[i].IsLess(shows[j])
	})
	return shows, nIdxes, nil
}
func (humanPlayer *HumanPlayer) DisplayOnlyHand() {
	indexLine, handsLine := humanPlayer.DisplayHand()
	humanPlayer.ioWriter.WriteString(fmt.Sprintf("%v\n", indexLine))
	humanPlayer.ioWriter.WriteString(fmt.Sprintf("%v\n", handsLine))
}
func (humanPlayer *HumanPlayer) DisplayName() {
	humanPlayer.ioWriter.WriteString(fmt.Sprintf("輪到%v了\n", humanPlayer.GetName()))
}
func (humanPlayer *HumanPlayer) InitPlay(topPlay []*Card, cardPatternHdr CardPatternHandlerInterface) []*Card {
	humanPlayer.DisplayName()
	isShowCorrect := false
	result := []*Card{}
	for !isShowCorrect {
		humanPlayer.DisplayOnlyHand()
		humanPlayer.ioWriter.Flush()
		cardsLine := humanPlayer.ShowCardStrategy.ShowCards(topPlay, humanPlayer.ioReader, humanPlayer.hands)
		isPass := strings.Compare(cardsLine, "-1") == 0
		if isPass && len(topPlay) == 0 {
			humanPlayer.ioWriter.WriteString("你不能在新的回合中喊 PASS\n")
			humanPlayer.ioWriter.Flush()
			isShowCorrect = false
			continue
		}
		shows, nIdxes, err := humanPlayer.ParesInputToShow(cardsLine)
		if err != nil {
			isShowCorrect = false
			continue
		}
		if strings.Compare(shows[0].String(), "C[3]") != 0 {
			humanPlayer.ioWriter.WriteString("此牌型不合法，請再嘗試一次。\n")
			humanPlayer.ioWriter.Flush()
			isShowCorrect = false
			continue
		}
		isShowCorrect = cardPatternHdr.DoCheckCardPattern(topPlay, shows, humanPlayer, humanPlayer.ioWriter)
		if isShowCorrect {
			humanPlayer.ExtractCards(nIdxes)
			result = shows
		}
		humanPlayer.ioWriter.Flush()
	}
	return result
}
