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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows the list of your todos",
	Long:  "List will show the list of your todos. You can filter the list with different options",
	Run:   runList,
}

func runList(cmd *cobra.Command, args []string) {
	// grab the items from the file
	items, err := todoList.ReadItems("C:/Users/bat20/.todos.json")

	// if there is an error, put it in the log
	if err != nil {
		log.Printf("%v", err)
	}

	// print out the item list
	fmt.Println(items)
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
