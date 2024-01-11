package todoList

import (
	"encoding/json"
	"os"
)

type Item struct {
	// text of the todo task
	Text string

	// priority of the todo task
	Priority int
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func SaveItems(filename string, items []Item) error {

	// get the items serialized, store that or the error that comes
	b, err := json.Marshal(items)
	// send back the error if there is one
	if err != nil {
		return err
	}

	// put the error message into a file
	err = os.WriteFile(filename, b, 0644)
	// send back the error if there is one
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	// try to read from the file
	b, err := os.ReadFile(filename)
	// give the error if there is one
	if err != nil {
		return []Item{}, err
	}
	// get the items, you have to declare the variable by type if you don't do :=
	var items []Item
	// make sure there is no problem in deserializing and deserialize with the pointer to the items variable
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	// give back the items with no error message
	return items, nil
}
