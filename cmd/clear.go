/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"todo/todoList"

	"github.com/spf13/cobra"
)

// variable for flags
var doneOnly bool

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all items",
	Long:  "Clear all of the tasks off of the todo list",
	Run:   runClear,
}

func runClear(cmd *cobra.Command, args []string) {
	if doneOnly {
		// get the items
		items, err := todoList.ReadItems(dataFile)
		if err != nil {
			log.Printf("%v", err)
		}

		// make a new list that only has the not done items
		var newItems []todoList.Item
		for _, i := range items {
			if !i.Done {
				newItems = append(newItems, i)
			}
		}

		// save this new list
		todoList.SaveItems(dataFile, newItems)
	} else {
		todoList.SaveItems(dataFile, nil)
	}
}

func init() {
	rootCmd.AddCommand(clearCmd)

	clearCmd.Flags().BoolVarP(&doneOnly, "done", "d", false, "Clear any items that are marked done from the Todo list")

}
