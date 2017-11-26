package wekan

import (
    "card"
    "list"
)

type Wekan struct {
    URL string
}

func (w Wekan) GetCards() ([]card.Card, error) {
    card_slice := make([]card.Card, 0)

    return card_slice, nil
}

func (w Wekan) GetLists() ([]list.List, error) {
    list_slice := make([]list.List, 0)

    return list_slice, nil
}

func (w Wekan) UpdateCard(c card.Card) error {
    return nil
}

func (w Wekan) NewCard(c card.Card) error {
    return nil
}
