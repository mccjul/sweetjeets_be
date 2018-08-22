package main

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
)

var db DB

// Resolver is the root resolver
type Resolver struct{}

// Hello resolver
func (r *Resolver) Hello() string { return "Hello, world!" }

// Bye resolver
func (r *Resolver) Bye(ctx context.Context, args *struct {
	Thingy string
}) string {
	return args.Thingy
}

// GetTodos resolver
func (r *Resolver) GetTodos(ctx context.Context) *[]*Todo {
	var resolvers = make([]*Todo, 0, len(todos))
	for _, t := range todos {
		resolvers = append(resolvers, &t)
	}

	return &resolvers
}

// AddTodo resolver
func (r *Resolver) AddTodo(ctx context.Context, args struct{ TodoIn TodoInput }) *Todo {
	return &Todo{Name: args.TodoIn.Name, Completed: false}
}

// UpdateTodo resolver
func (r *Resolver) UpdateTodo(ctx context.Context, args struct{ TodoIn TodoInput }) *Todo {
	return &Todo{Name: args.TodoIn.Name, Completed: args.TodoIn.Completed}
}

// DeleteTodo resolver
func (r *Resolver) DeleteTodo(ctx context.Context, args *struct{ ID int32 }) *Todo {
	return &Todo{Name: "rex", Completed: true}
}

func gqlIDP(id uint) *graphql.ID {
	r := graphql.ID(fmt.Sprint(id))
	return &r
}

var todos = []Todo{
	Todo{Name: "rex", Completed: true},
	Todo{Name: "goldie", Completed: true},
	Todo{Name: "spot", Completed: true},
	Todo{Name: "pokey", Completed: true},
	Todo{Name: "sneezy", Completed: false},
	Todo{Name: "duke", Completed: true},
	Todo{Name: "duchess", Completed: false},
	Todo{Name: "bernard", Completed: true},
	Todo{Name: "William III of Chesterfield", Completed: true},
	Todo{Name: "hops", Completed: true},
}
