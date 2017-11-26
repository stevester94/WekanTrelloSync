package card
import "fmt"

type Card struct {
    ID string
    ParentListID string
    Name string
    DateLastActivity string
}

func (c Card) Print() {
    fmt.Println("   Name:", c.Name)
    fmt.Println("   ID:", c.ID)
    fmt.Println("   DateLastActivity:", c.DateLastActivity)
    fmt.Println("   ParentListID:", c.ParentListID)
}