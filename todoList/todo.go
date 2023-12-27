package todoList

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(string(b))

	return nil
}
