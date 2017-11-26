package main


import (
    "board"
    "card"
    "list"
    "fmt"
    "wekan"
    "trello"
    "strings"
)


func main() {
    fmt.Println("Wekan exercise!")

    my_wekan := wekan.Wekan{"192.168.1.220"}

    var my_cards []card.Card
    my_cards, _ = my_wekan.GetCards()

    var my_list []list.List
    my_list, _ = my_wekan.GetLists()

    fmt.Println("my_list:", my_list)
    fmt.Println("my_cards:", my_cards)

    var b board.Board = my_wekan
    fmt.Println("Board:", b)



    fmt.Println("\nTrello exercise!")

    my_trello := trello.Trello{"trello.com"}

    my_cards, _ = my_trello.GetCards()
    my_list, _ = my_trello.GetLists()

    fmt.Println("my_list(trello):", my_list)

    fmt.Println("All Trello cards:")
    for _, c := range my_cards {
        c.Print()
        fmt.Println()
    }

    fmt.Println("Update test...")
    var target_card card.Card
    for _, c := range my_cards {
        if strings.Contains(c.Name, "LETS GO") {
            target_card = c
        }
    }

    target_card.Name = target_card.Name + "O"
    my_trello.UpdateCard(target_card)
    my_cards, _ = my_trello.GetCards()


    fmt.Println("All Trello cards after update:")
    for _, c := range my_cards {
        c.Print()
        fmt.Println()
    }

    b = my_trello
    fmt.Println("Board:", b)
}
