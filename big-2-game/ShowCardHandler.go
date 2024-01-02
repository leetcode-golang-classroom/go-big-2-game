package big_2_game

type ShowCardHandlerInterface interface {
	SetNext(next ShowCardHandlerInterface)
	DoShowCards(topPlay []*Card, hand []*Card) string
	IsMatch(topPlay []*Card, hand []*Card) bool
	ShowCards(topPlay []*Card, hand []*Card) string
}

type ShowCardHandler struct {
	ShowCardHandlerInterface
	Next ShowCardHandlerInterface
}

func NewShowCardsHandler(showCardsHdr ShowCardHandlerInterface) ShowCardHandlerInterface {
	return &ShowCardHandler{
		showCardsHdr,
		nil,
	}
}
func (showCardHdr *ShowCardHandler) SetNext(next ShowCardHandlerInterface) {
	showCardHdr.Next = next
}

func (showCardHdr *ShowCardHandler) DoShowCards(topPlay []*Card, hand []*Card) string {
	initHdr := showCardHdr
	for initHdr != nil {
		if initHdr.IsMatch(topPlay, hand) {
			return initHdr.ShowCards(topPlay, hand)
		} else {
			if initHdr.Next == nil {
				initHdr = nil
			} else {
				initHdr = (initHdr.Next).(*ShowCardHandler)
			}
		}
	}
	return "-1"
}
