/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"todo/todoList"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the list",
	Long:  `Add will create a new item on the todo list`,
	Run:   runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	// declare the list of items (new or existing)
	items, err := todoList.ReadItems("C:/Users/bat20/.todos.json")
	// if there was an error, tell us
	if err != nil {
		log.Printf("%v", err)
	}

	// add the items to the list
	for _, x := range args {
		// Text is the variable name in the struct, you "construct" the Item with the x value by the :
		items = append(items, todoList.Item{Text: x})
	}
	// save the list
	err = todoList.SaveItems("C:/Users/bat20/.todos.json", items)

	// if there was an error, tell us
	if err != nil {
		log.Printf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
