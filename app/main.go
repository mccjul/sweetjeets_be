package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (q *query) Hello() string { return "Hello, world!" }

func (q *query) Bye(ctx context.Context, args *struct {
	Thingy string
}) string {
	return args.Thingy
}

func getFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func main() {
	s, err := getFile("./schema.graphql")
	page, err := getFile("./graphql.html")
	if err != nil {
		return
	}

	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(page))
	}))

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev modes")
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
