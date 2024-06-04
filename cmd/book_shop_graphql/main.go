package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

type Book struct {
	ID            uint64    `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	DateOfRealise time.Time `json:"date_of_realise"`
}

var books []Book

var bookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id":              &graphql.Field{Type: graphql.Int},
			"title":           &graphql.Field{Type: graphql.String},
			"author":          &graphql.Field{Type: graphql.String},
			"date_of_realise": &graphql.Field{Type: graphql.DateTime},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Type:        bookType,
				Description: "get book by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						for _, book := range books {
							if int(book.ID) == id {
								return book, nil
							}

						}
					}
					return nil, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(bookType),
				Description: "Get books list",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return books, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        bookType,
			Description: "Create new book",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"author": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"date_of_reliase": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				rand.Seed(time.Now().UnixNano())
				book := Book{
					ID:            uint64(rand.Intn(100000)),
					Title:         params.Args["title"].(string),
					Author:        params.Args["author"].(string),
					DateOfRealise: params.Args["date_of_realise"].(time.Time),
				}
				books = append(books, book)
				return books, nil
			},
		},
		"update": &graphql.Field{
			Type:        bookType,
			Description: "Update book by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"author": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"date_of_reliase": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				title, titleOk := params.Args["title"].(string)
				author, authorOk := params.Args["author"].(string)
				data_of_realise, data_of_realiseOk := params.Args["data_of_realise"].(time.Time)
				book := Book{}
				for i, b := range books {
					if uint64(id) == b.ID {
						if titleOk {
							books[i].Title = title
						}
						if authorOk {
							books[i].Author = author
						}
						if data_of_realiseOk {
							books[i].DateOfRealise = data_of_realise
						}
						book = books[i]
						break

					}
				}
				return book, nil

			},
		},
		"delete": &graphql.Field{
			Type:        bookType,
			Description: "Delete book by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(int)
				book := Book{}
				for i, b := range books {
					if uint64(id) == b.ID {
						book = books[i]
						books = append(books[:i], books[i+1:]...)
					}
				}
				return book, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func initBooksData(b *[]Book) {
	book1 := Book{ID: 1, Title: "Война и мир", Author: "Л.Н. Толстой", DateOfRealise: time.Date(1869, 0, 0, 0, 0, 0, 0, &time.Location{})}
	book2 := Book{ID: 2, Title: "Финансист", Author: "Теодор Дрейзер", DateOfRealise: time.Date(1912, 0, 0, 0, 0, 0, 0, &time.Location{})}
	book3 := Book{ID: 3, Title: "100 Go mistakes and how avoid them", Author: "TEIVA HARSANYI", DateOfRealise: time.Date(2022, 0, 0, 0, 0, 0, 0, &time.Location{})}
	*b = append(*b, book1, book2, book3)
}

func main() {
	initBooksData(&books)
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)

		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
