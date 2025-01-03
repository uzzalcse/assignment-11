package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type ResultSet struct {
    UserID int    `json:"userId"`
    PostID int    `json:"id"`
    Title string `json:"title"`
    Body string `json:"body"`
}

func main() {

    response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    if err != nil {
        fmt.Printf("error making HTTP request: %v\n", err)
        return
    }
	
    defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("error reading response body: %v\n", err)
        return
    }

	var resultSet1 ResultSet
    err = json.Unmarshal(body, &resultSet1)

    if err != nil {
        fmt.Printf("error unmarshaling JSON: %v\n", err)
        return
    }

    fmt.Printf("User ID: %d\n", resultSet1.UserID)
    fmt.Printf("Post ID: %d\n", resultSet1.PostID)
    fmt.Printf("Title: %s\n", resultSet1.Title)
    fmt.Printf("Body: %s\n", resultSet1.Body)
}