package todoList

import (
	"encoding/json"
	"os"
	"strconv"
)

type Item struct {
	// text of the todo task
	Text string

	// priority of the todo task
	Priority int

	// position in the list for sorting purposes
	Position int

	// status of the item
	Done bool
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	case 4:
		i.Priority = 4
	default:
		i.Priority = 2
	}
}

func (i *Item) Label() string {
	return strconv.Itoa(i.Position) + "."
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	} else {
		return ""
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
	// set their positions
	for i, _ := range items {
		items[i].Position = i + 1
	}
	// give back the items with no error message
	return items, nil
}

// ByPri uses sort.Interface for []Item based on Priority and Position
type ByPri []Item

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	// sort based on done-ness
	if s[i].Done != s[j].Done {
		return !s[i].Done
	}
	// then on priority
	if s[i].Priority == s[j].Priority {
		// then on position
		return s[i].Position < s[j].Position
	}
	return s[i].Priority < s[j].Priority
}
