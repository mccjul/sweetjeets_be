package main

import (
	"log"

	"github.com/jinzhu/gorm"
	// nolint: gotype
	_ "github.com/go-sql-driver/mysql"
)

// DB something someting
type DB struct {
	DB *gorm.DB
}

func newDB() (*DB, error) {
	// connect to the example db, create if it doesn't exist
	db, err := gorm.Open("mysql", "root:root@tcp(172.17.0.2:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}

	// drop tables and all data, and recreate them fresh for this run
	db.DropTableIfExists(&Todo{})
	db.AutoMigrate(&Todo{})

	// put all the users into the db
	for _, t := range todos {
		if err := db.Create(&t).Error; err != nil {
			return nil, err
		}
	}

	return &DB{db}, nil
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
