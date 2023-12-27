package todoList

import (
	"encoding/json"
	"os"
)

type Item struct {
	Text string
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
