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
	ID        *float64
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
func (db *DB) getTodos(ctx context.Context) (*[]*Todo, error) {
	var t []*Todo
	err := db.DB.Find(&t).Error
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (db *DB) addTodo(ctx context.Context, t *TodoInput) (*Todo, error) {
	todo := Todo{
		Name:      t.Name,
		Completed: t.Completed,
	}

	err := db.DB.Create(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (db *DB) updateTodo(ctx context.Context, tin *TodoInput) (*Todo, error) {
	var t Todo

	err := db.DB.First(&t, int(*tin.ID)).Error
	if err != nil {
		return nil, err
	}

	updated := Todo{
		Name:      tin.Name,
		Completed: tin.Completed,
	}

	err = db.DB.Model(&t).Updates(updated).Error
	if err != nil {
		return nil, err
	}

	err = db.DB.First(&t, int(*tin.ID)).Error
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (db *DB) deleteTodo(ctx context.Context, id *int32) (*Todo, error) {
	var t Todo
	err := db.DB.First(&t, int(*id)).Error
	if err != nil {
		return nil, err
	}

	err = db.DB.Delete(&t).Error
	if err != nil {
		return nil, err
	}

	return &t, nil
}
