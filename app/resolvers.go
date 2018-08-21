package main

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
)

var db DB

// Resolver is the root resolver
type Resolver struct{}

func (r *Resolver) Hello() string { return "Hello, world!" }

func (r *Resolver) Bye(ctx context.Context, args *struct {
	Thingy string
}) string {
	return args.Thingy
}

func (r *Resolver) GetTodos(ctx context.Context) *Todo {
	return &Todo{Name: "rex", Completed: true}
}

func (r *Resolver) AddTodo(ctx context.Context, t *Todo) *Todo {
	return t
}

func (r *Resolver) UpdateTodo(ctx context.Context, t *Todo) *Todo {
	return t
}

func (r *Resolver) DeleteTodo(ctx context.Context, id int) *Todo {
	return &Todo{Name: "rex", Completed: true}
}

func gqlIDP(id uint) *graphql.ID {
	r := graphql.ID(fmt.Sprint(id))
	return &r
}
