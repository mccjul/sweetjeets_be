package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Name      string
	Completed bool
}

// RESOLVERS ===========================================================================
//ID resolves the ID field for Pet
func (t *Todo) ID(ctx context.Context) *graphql.ID {
	return gqlIDP(t.Model.ID)
}

// NAME resolves the name field for Pet
func (t *Todo) NAME(ctx context.Context) *string {
	return &t.Name
}

func (t *Todo) COMPLETED(ctx context.Context) *bool {
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

// var todos = []Todo{
// 	Todo{name: "rex", id: 1, completed: true},
// 	Todo{name: "goldie", id: 2, completed: true},
// 	Todo{name: "spot", id: 3, completed: true},
// 	Todo{name: "pokey", id: 4, completed: true},
// 	Todo{name: "sneezy", id: 5, completed: false},
// 	Todo{name: "duke", id: 6, completed: true},
// 	Todo{name: "duchess", id: 7, completed: false},
// 	Todo{name: "bernard", id: 8, completed: true},
// 	Todo{name: "William III of Chesterfield", id: 9, completed: true},
// 	Todo{name: "hops", id: 10, completed: true},
// }

// func (db *DB) getTodos(ctx context.Context) *[]Todo {
// 	return &todos
// }

// func (db *DB) updateTodo(ctx context.Context, t Todo) Todo {
// 	return t
// }

// func (db *DB) deleteTodo(ctx context.Context, id int32) Todo {
// 	return todos[id]
// }
