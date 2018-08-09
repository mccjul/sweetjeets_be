package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

// Todo is the base type for pets to be used by the db and gql
type Todo struct {
	gorm.Model
	OwnerID uint
	Name    string
	Tags    []Tag `gorm:"many2many:pet_tags"`
}

// RESOLVERS ===========================================================================
// ID resolves the ID field for Pet
func (t *Todo) ID(ctx context.Context) *graphql.ID {
	return gqlIDP(t.Model.ID)
}

// OWNER resolves the owner field for Pet
func (t *Todo) OWNER(ctx context.Context) (*User, error) {
	return db.getPetOwner(ctx, int32(t.OwnerID))
}

// NAME resolves the name field for Pet
func (t *Todo) NAME(ctx context.Context) *string {
	return &t.Name
}

// TAGS resolves the pet tags
func (t *Todo) TAGS(ctx context.Context) (*[]*Tag, error) {
	return db.getPetTags(ctx, t)
}

// DB ===================================================================================
// GetPet should authorize the user in ctx and return a pet or error
func (db *DB) getPet(ctx context.Context, id int32) (*Pet, error) {
	var p Pet
	err := db.DB.First(&p, id).Error
	if err != nil {
		return nil, err
	}

	return &p, nil
}
