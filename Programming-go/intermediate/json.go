package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Active   bool     `json:"is_active"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// Marshal (struct to JSON)
	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
		Active:   true,
		Tags:     []string{"admin", "developer"},
	}

	fmt.Printf("User created: %+v\n", user)

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Println("JSON:", string(jsonData))

	// Marshal with indent
	prettyJSON, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("Pretty JSON:\n", string(prettyJSON))

	// Unmarshal (JSON to struct)
	jsonStr := `{"id":2,"username":"jane_doe","email":"jane@example.com","is_active":true}`
	var user2 User
	err = json.Unmarshal([]byte(jsonStr), &user2)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	fmt.Printf("Unmarshaled: %+v\n", user2)
}