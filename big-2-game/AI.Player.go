package big_2_game

import (
	"bufio"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

type AIPlayer struct {
	*Player
	ioWriter *bufio.Writer
}

var AiPlayerNames map[string]struct{} = make(map[string]struct{})

func NewAIPlayer(ioWriter *bufio.Writer, showCardStrategy ShowCardStrategyInterface) *AIPlayer {
	return &AIPlayer{
		Player:   NewPlayer(showCardStrategy),
		ioWriter: ioWriter,
	}
}
func (aiPlayer *AIPlayer) RandomString(strlen int) string {
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[randGenerator.Intn(len(chars))]
	}
	return string(result)
}
func (aiPlayer *AIPlayer) NameSelf() {
	var name string
	nameExists := true
	for nameExists {
		name = aiPlayer.RandomString(10)
		_, nameExists = AiPlayerNames[name]
		if !nameExists {
			AiPlayerNames[name] = struct{}{}
		}
	}
	aiPlayer.SetName(name)
}

func (aiPlayer *AIPlayer) Play(topPlay []*Card, cardPatternHdr CardPatternHandlerInterface) []*Card {
	isShowCorrect := false
	aiPlayer.DisplayName()
	result := []*Card{}
	for !isShowCorrect {
		aiPlayer.DisplayOnlyHand()
		aiPlayer.ioWriter.Flush()
		cardsLine := aiPlayer.ShowCardStrategy.ShowCards(topPlay, nil, aiPlayer.hands)
		isPass := strings.Compare(cardsLine, "-1") == 0
		if isPass && len(topPlay) != 0 {
			aiPlayer.ioWriter.WriteString(fmt.Sprintf("玩家 %v PASS.\n", aiPlayer.GetName()))
			aiPlayer.ioWriter.Flush()
			isShowCorrect = true
			return []*Card{}
		}
		if isPass && len(topPlay) == 0 {
			aiPlayer.ioWriter.WriteString("你不能在新的回合中喊 PASS\n")
			aiPlayer.ioWriter.Flush()
			isShowCorrect = false
			continue
		}
		shows, nIdxes, err := aiPlayer.ParesInputToShow(cardsLine)
		if err != nil {
			isShowCorrect = false
			continue
		}
		isShowCorrect = cardPatternHdr.DoCheckCardPattern(topPlay, shows, aiPlayer, aiPlayer.ioWriter)
		if isShowCorrect {
			aiPlayer.ExtractCards(nIdxes)
			result = shows
		}
		aiPlayer.ioWriter.Flush()
	}
	return result
}

func (aiPlayer *AIPlayer) InitPlay(topPlay []*Card, cardPatternHdr CardPatternHandlerInterface) []*Card {
	isShowCorrect := false
	aiPlayer.DisplayName()
	result := []*Card{}
	for !isShowCorrect {
		aiPlayer.DisplayOnlyHand()
		aiPlayer.ioWriter.Flush()
		// cardsLine := "0"
		cardsLine := aiPlayer.InitChecker(aiPlayer.hands)
		isPass := strings.Compare(cardsLine, "-1") == 0
		if isPass && len(topPlay) != 0 {
			aiPlayer.ioWriter.WriteString(fmt.Sprintf("玩家 %v PASS.\n", aiPlayer.GetName()))
			aiPlayer.ioWriter.Flush()
			isShowCorrect = true
			return []*Card{}
		}
		if isPass && len(topPlay) == 0 {
			aiPlayer.ioWriter.WriteString("你不能在新的回合中喊 PASS\n")
			aiPlayer.ioWriter.Flush()
			isShowCorrect = false
			continue
		}
		shows, nIdxes, err := aiPlayer.ParesInputToShow(cardsLine)
		if err != nil {
			fmt.Println(err)
			isShowCorrect = false
			continue
		}
		isShowCorrect = cardPatternHdr.DoCheckCardPattern(topPlay, shows, aiPlayer, aiPlayer.ioWriter)
		if isShowCorrect {
			aiPlayer.ExtractCards(nIdxes)
			result = shows
		}
		aiPlayer.ioWriter.Flush()
	}
	return result
}

func (aiPlayer *AIPlayer) ParesInputToShow(cardsLine string) ([]*Card, []int, error) {
	idxes := strings.Split(cardsLine, " ")
	shows := []*Card{}
	nIdxes := []int{}
	for _, idx := range idxes {
		num, err := strconv.Atoi(idx)
		if err != nil {
			return nil, nil, err
		}
		if num < -1 || num > len(aiPlayer.hands)-1 {
			return nil, nil, fmt.Errorf("%v out of index range", num)
		}
		nIdxes = append(nIdxes, num)
		shows = append(shows, aiPlayer.hands[num])
	}
	sort.Slice(shows, func(i, j int) bool {
		return shows[i].IsLess(shows[j])
	})
	return shows, nIdxes, nil
}
func (aiPlayer *AIPlayer) DisplayOnlyHand() {
	indexLine, handsLine := aiPlayer.DisplayHand()
	aiPlayer.ioWriter.WriteString(fmt.Sprintf("%v\n", indexLine))
	aiPlayer.ioWriter.WriteString(fmt.Sprintf("%v\n", handsLine))
}
func (aiPlayer *AIPlayer) DisplayName() {
	aiPlayer.ioWriter.WriteString(fmt.Sprintf("輪到%v了\n", aiPlayer.GetName()))
}
func (aiPlayer *AIPlayer) InitChecker(hand []*Card) string {
	cardLines := "0"
	rankMap := make(map[int][]int)
	for idx, card := range hand {
		list, ok := rankMap[int(card.rank)]
		if !ok {
			rankMap[int(card.rank)] = []int{idx}
		} else {
			rankMap[int(card.rank)] = append(list, idx)
		}
	}
	startRank := 3
	result := []int{}
	for idx := 1; idx <= 4; idx++ {
		list, ok := rankMap[startRank+idx]
		if !ok {
			break
		} else {
			result = append(result, list[0])
		}
	}
	if len(result) == 4 {
		return fmt.Sprintf("0 %v %v %v %v", result[0], result[1], result[2], result[3])
	}
	list := rankMap[3]
	remain := 5 - len(list)
	following := []int{}
	if len(list) == 2 || len(list) == 3 {
		for _, remainList := range rankMap {
			if len(remainList) == remain {
				following = remainList
				break
			}
		}
	}
	if len(following) != 0 {
		result := ""
		for idx, card := range list {
			result += fmt.Sprintf("%v", card)
			if idx != len(list)-1 {
				result += " "
			}
		}
		for _, card := range following {
			result += fmt.Sprintf(" %v", card)
		}
		return result
	}
	if len(list) >= 2 {
		return fmt.Sprintf("%v %v", list[0], list[1])
	}
	return cardLines
}
