package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

// Todo datatype
type Todo struct {
	gorm.Model
	Name      string
	Completed bool
}

// TodoInput datatype
type TodoInput struct {
	Name      string
	Completed bool
}

//ID resolves the ID field for Pet
func (t *Todo) ID(ctx context.Context) *graphql.ID {
	return gqlIDP(t.Model.ID)
}

// NAME resolves the name field for Pet
func (t *Todo) NAME(ctx context.Context) *string {
	return &t.Name
}

// COMPLETED resolves the name field for Todo
func (t *Todo) COMPLETED(ctx context.Context) *bool {
	return &t.Completed
}

// NAME resolves the name field for Pet
func (t *TodoInput) NAME(ctx context.Context) *string {
	return &t.Name
}

// COMPLETED resolves the name field for Todo
func (t *TodoInput) COMPLETED(ctx context.Context) *bool {
	return &t.Completed
}

// DB ===================================================================================
// GetPet should authorize the user in ctx and return a pet or error
// func (db *DB) getPet(ctx context.Context, id int32) (*Pet, error) {
// 	var p Pet
// 	err := db.DB.First(&p, id).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &p, nil
// }

func (db *DB) getTodos(ctx context.Context) *[]*Todo {
	var resolvers = make([]*Todo, 0, len(todos))
	for _, t := range todos {
		resolvers = append(resolvers, &t)
	}

	return &resolvers
}

func (db *DB) addTodo(ctx context.Context, t *TodoInput) *Todo {
	return &Todo{Name: t.Name, Completed: false}
}

func (db *DB) updateTodo(ctx context.Context, t *TodoInput) *Todo {
	return &Todo{Name: t.Name, Completed: t.Completed}
}

func (db *DB) deleteTodo(ctx context.Context, id *int32) *Todo {
	return &Todo{Name: "rex", Completed: true}
}
