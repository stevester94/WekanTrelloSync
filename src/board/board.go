package board

import (
    "card"
    "list"
)


type Board interface {
    GetCards() ([]card.Card, error)
    GetLists() ([]list.List, error)
    UpdateCard(c card.Card) error
    NewCard(c card.Card) error
}
