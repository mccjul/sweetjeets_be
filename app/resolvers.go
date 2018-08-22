package main

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
)

var db DB

func init() {
	// var err error
	// d, err := newDB("./db.sqlite")
	// if err != nil {
	// 	panic(err)
	// }

	// db = *d
}

// Resolver is the root resolver
type Resolver struct{}

// GetTodos resolver
func (r *Resolver) GetTodos(ctx context.Context) (*[]*Todo, error) {
	return db.getTodos(ctx)
}

// AddTodo resolver
func (r *Resolver) AddTodo(ctx context.Context, args struct{ TodoIn TodoInput }) (*Todo, error) {
	return db.addTodo(ctx, &args.TodoIn)
}

// UpdateTodo resolver
func (r *Resolver) UpdateTodo(ctx context.Context, args struct{ TodoIn TodoInput }) (*Todo, error) {
	return db.updateTodo(ctx, &args.TodoIn)
}

// DeleteTodo resolver
func (r *Resolver) DeleteTodo(ctx context.Context, args *struct{ ID int32 }) (*Todo, error) {
	return db.deleteTodo(ctx, &args.ID)
}

func gqlIDP(id uint) *graphql.ID {
	r := graphql.ID(fmt.Sprint(id))
	return &r
}
