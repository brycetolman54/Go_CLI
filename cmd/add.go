/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"todo/todoList"

	"github.com/spf13/cobra"
)

// priority value for adding a task
var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the list",
	Long:  `Add will create a new item on the todo list`,
	Run:   runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	// make sure there are args
	if len(args) == 0 {
		fmt.Println("you have to include a task to add to the list")
		return
	}

	// declare the list of items (new or existing)
	items, err := todoList.ReadItems(dataFile)
	// if there was an error, tell us
	if err != nil {
		log.Printf("%v", err)
	}

	// add the items to the list
	for _, x := range args {
		// get the item
		item := todoList.Item{Text: x}
		// set the priority
		item.SetPriority(priority)
		// Text is the variable name in the struct, you "construct" the Item with the x value by the :
		items = append(items, item)
	}
	// save the list
	err = todoList.SaveItems(dataFile, items)

	// if there was an error, tell us
	if err != nil {
		log.Printf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1, 2, 3, 4")

}
