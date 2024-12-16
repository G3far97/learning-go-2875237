package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://jsonplaceholder.typicode.com/todos"

func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	// fmt.Print(content)

	users := usersFromJson(content)
	for _, user := range users {
		fmt.Println(user.Title)
	}
}

func usersFromJson(content string) []User {
	users := make([]User, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var user User
	for decoder.More() {
		err := decoder.Decode(&user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}

type User struct {
	ID    int
	Title string
}
