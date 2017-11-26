package main

import "fmt"
import "net/http"
import "io/ioutil"
//import "os"
import "encoding/json"
import "bytes"

const Base_URL string = "http://192.168.1.220:80801"

const Username string = "steven"
const Password string = "test123"

func FetchToken() (string) {
    endpoint_format_string string = "/users/login/"
    request_url := Base_URL+endpoint_format_string
    fmt.Printf("request_url: %s\n", request_url)


    payload_map := make(map[string]interface{})
    var payload_bytes  []byte

    payload_map["username"] = Username
    payload_map["password"] = Password

    payload_bytes, _ = json.Marshal(payload_map)

    fmt.Printf("Payload Bytes: \n%s\n", string(payload_bytes))


    client := &http.Client{} // No fucking clue what this is really doing

    req, err := http.NewRequest(http.MethodPost, request_url, bytes.NewBuffer(payload_bytes))
    if err != nil {
        panic(err)
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("There was an error when attempting to make the request...")
        panic(err)
    }


    defer resp.Body.Close() // Believe defer is some kinf of synchronization thing

    contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Response contents:\n%s\n", string(contents))

    response_map := make(map[string]string)

    json.Unmarshal(contents, &response_map)
    var token string = response_map["token"]
    fmt.Printf("Token: %s\n", token)

    return token
}

func fetchBoards(token string) ([]map[string]string) {
    

func main() {
    token := FetchToken()

    fmt.Println("Have token:", token)
}


















