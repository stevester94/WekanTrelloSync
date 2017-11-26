package trello

import (
    "card"
    "list"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"
	"fmt"
    "bytes"
)

const API_Key string = "f44bfd10a8cfed257d7dc6cb1ba18038"
const Token string = "90f6221553f46fb8bcedd905e3f74cd7e13327336c7f18ec6754620bf520627f"
const Board_ID string = "db6BDqSJ"
const Board_endpoint_format_string string = "https://api.trello.com/1/boards/%s/cards?key=%s&token=%s" // Board_ID, key, token
const Card_endpoint_format_string string = "https://api.trello.com/1/cards/%s?key=%s&token=%s" // Card_ID, key, token


type Trello struct {
    URL string
}

func (t Trello) GetCards() ([]card.Card, error) {
    request_url := fmt.Sprintf(Board_endpoint_format_string, Board_ID, API_Key, Token)

	resp, err := http.Get(request_url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close() // Believe defer is some kind of synchronization thing

	contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
	    fmt.Printf("%s", err)
	    os.Exit(1)
    }


    cards := deJsonifyCardArray(contents)
    return cards, nil
}

func (t Trello) GetLists() ([]list.List, error) {
    list_slice := make([]list.List, 0)

    return list_slice, nil
}

func (t Trello) UpdateCard(c card.Card) error {
    request_url := fmt.Sprintf(Card_endpoint_format_string, c.ID, API_Key, Token)
    fmt.Printf("request_url: %s\n\n", request_url)

    payload_map := make(map[string]interface{})
    var payload_bytes []byte

    // Only updating card names for now!
    payload_map["name"] = c.Name


    payload_bytes, _ = json.Marshal(payload_map)
    fmt.Printf("payload_bytes: %s\n\n", payload_bytes)



    // PUT is not easily accessed! Need to use the client interface of http lib...
    client := &http.Client{}


    req, err := http.NewRequest(http.MethodPut, request_url, bytes.NewBuffer(payload_bytes))
    if err != nil {
        panic(err)
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }


    defer resp.Body.Close() // Believe defer is some kind of synchronization thing

    return nil
}

func (t Trello) NewCard(c card.Card) error {
    return nil
}




// We expect to get a slice of maps from the input JSON
// Each element is a card
// We are getting cards because that's which endpoint we requested
func deJsonifyCardArray(s []byte) []card.Card {
	var dat = make([]map[string]interface{}, 50)
	return_cards := make([]card.Card, 0)

	if err := json.Unmarshal(s, &dat); err != nil {
        panic(err)
    }

    for _, card_map := range dat {
    	var c card.Card

        // Populate our card object and append it to the return slice
        c.ID = card_map["id"].(string)
        c.Name = card_map["name"].(string)
        c.DateLastActivity = card_map["dateLastActivity"].(string)
        c.ParentListID = card_map["idList"].(string)

        return_cards = append(return_cards, c)
    }

    return return_cards
}