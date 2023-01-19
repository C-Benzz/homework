package graph

import (
	"fmt"

	"github.com/C-Benzz/assignment1.git/graph/model"
	"github.com/google/uuid"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Database struct {
	TodoTable map[string]model.Todo
}

func (db *Database) AddNewTodo(input *model.Todo) error {
	newID := uuid.New().String()
	input.ID = newID
	db.TodoTable[newID] = *input
	return nil
}

func (db *Database) UpdateTodoByID(input model.Todo) error {
	if _, ok := db.TodoTable[input.ID]; !ok {
		return fmt.Errorf("movie id: %s was not found", input.ID)
	}
	db.TodoTable[input.ID] = input
	return nil
}

func (db *Database) DeleteTodoByID(ID string) error {
	if _, ok := db.TodoTable[ID]; !ok {
		return fmt.Errorf("movie id: %s was not found", ID)
	}
	delete(db.TodoTable, ID)
	return nil
}

func (db *Database) GetTodoByID(ID string) (*model.Todo, error) {
	if ret, ok := db.TodoTable[ID]; ok {
		return &ret, nil
	} else {
		return nil, fmt.Errorf("movie id: %s was not found", ID)
	}
}
func (db *Database) ListAllTodo() []*model.Todo {
	listTodo := []*model.Todo{}
	for _, n := range db.TodoTable {
		newTodo := n
		listTodo = append(listTodo, &newTodo)
	}
	return listTodo
}

type Resolver struct {
	DB Database
}
