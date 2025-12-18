package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Book represents a book
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

// User represents a user
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// Sample data
var books = []Book{
	{ID: 1, Title: "The Go Programming Language", Author: "Alan Donovan", Year: 2015},
	{ID: 2, Title: "Clean Code", Author: "Robert C. Martin", Year: 2008},
	{ID: 3, Title: "Design Patterns", Author: "Gang of Four", Year: 1994},
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
	{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
	{ID: 3, Name: "Charlie", Email: "charlie@example.com", Age: 35},
}

// openBrowser opens the specified URL in the default browser
func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	}
	if err != nil {
		fmt.Printf("Failed to open browser: %v\n", err)
	}
}

func main() {
	fmt.Println("Starting GraphQL server on port 8081...")

	// Define Book type
	bookType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id":     &graphql.Field{Type: graphql.Int},
			"title":  &graphql.Field{Type: graphql.String},
			"author": &graphql.Field{Type: graphql.String},
			"year":   &graphql.Field{Type: graphql.Int},
		},
	})

	// Define User type
	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.Int},
			"name":  &graphql.Field{Type: graphql.String},
			"email": &graphql.Field{Type: graphql.String},
			"age":   &graphql.Field{Type: graphql.Int},
		},
	})

	// Define Root Query
	fields := graphql.Fields{
		// Simple hello query
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},

		// Get all books
		"books": &graphql.Field{
			Type: graphql.NewList(bookType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return books, nil
			},
		},

		// Get book by ID
		"book": &graphql.Field{
			Type: bookType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(int)
				for _, book := range books {
					if book.ID == id {
						return book, nil
					}
				}
				return nil, nil
			},
		},

		// Get all users
		"users": &graphql.Field{
			Type: graphql.NewList(userType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return users, nil
			},
		},

		// Get user by ID
		"user": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(int)
				for _, user := range users {
					if user.ID == id {
						return user, nil
					}
				}
				return nil, nil
			},
		},

		// Search books by author
		"booksByAuthor": &graphql.Field{
			Type: graphql.NewList(bookType),
			Args: graphql.FieldConfigArgument{
				"author": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				author := p.Args["author"].(string)
				var result []Book
				for _, book := range books {
					if book.Author == author {
						result = append(result, book)
					}
				}
				return result, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(err)
	}

	// Create a graphql-go/handler
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	r := gin.Default()

	// Serve GraphQL API
	r.POST("/graphql", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	// Serve GraphiQL interface
	r.GET("/graphql", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	// Auto open GraphiQL in browser
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\n🚀 Opening GraphiQL at http://localhost:8081/graphql")
		fmt.Println("\n📝 Example Queries:")
		fmt.Println("   { hello }")
		fmt.Println("   { books { id title author year } }")
		fmt.Println("   { book(id: 1) { title author } }")
		fmt.Println("   { users { id name email } }")
		fmt.Println("   { user(id: 1) { name email age } }")
		fmt.Println("   { booksByAuthor(author: \"Alan Donovan\") { title year } }")
		openBrowser("http://localhost:8081/graphql")
	}()

	r.Run(":8081") // Running on 8081 to avoid conflict with REST example
}
