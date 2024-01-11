/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"todo/todoList"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark an item on the list as done",
	Long:    "Mark will mark an item on the list as done and will remove it from the list",
	Run:     runDone,
}

func runDone(cmd *cobra.Command, args []string) {
	// make sure there is an argument there
	if len(args) == 0 {
		fmt.Println("you have to provide a label for the task to mark as done")
		return
	}

	// get the items
	items, err := todoList.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}

	// get the label/position
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	// make sure the item is on the list
	if i > 0 && i <= len(items) {
		// mark it done
		items[i-1].Done = true
		// tell them you marked it done
		fmt.Printf("%q %v\n", items[i-1].Text, "marked done!")

		// resort and reassign positions on the list
		sort.Sort(todoList.ByPri(items))
		todoList.SaveItems(dataFile, items)
	} else {
		log.Println(i, "doesn't match any items on the list")
	}
}

func init() {
	rootCmd.AddCommand(doneCmd)

}
