/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
	"todo/todoList"

	"github.com/spf13/cobra"
)

// strings for coloring
var esc string = "\u001b"
var setText string = esc + "[38;5;"
var setBG string = esc + "[48;5;"
var redText string = setText + "160m"
var yellowText string = setText + "226m"
var greenText string = setText + "46m"
var blueText string = setText + "12m"
var restText string = setText + "15m"

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows the list of your todos",
	Long:  "List will show the list of your todos. You can filter the list with different options",
	Run:   runList,
}

func runList(cmd *cobra.Command, args []string) {
	// set up the writer
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	// grab the items from the file
	items, err := todoList.ReadItems(dataFile)

	// sort the items first by priority and then by position
	sort.Sort(todoList.ByPri(items))

	// if there is an error, put it in the log
	if err != nil {
		log.Printf("%v", err)
	}

	// add the items to the list and send it
	for _, i := range items {
		var color string
		var extra string
		// get the priority
		pt := i.Priority
		// set the text color
		switch pt {
		case 1:
			color = redText
		case 2:
			color = yellowText
		case 3:
			color = greenText
		case 4:
			color = blueText
		default:
			color = ""
		}
		if pt == 1 || pt == 2 {
			extra = " "
		} else {
			extra = ""
		}
		fmt.Fprintln(w, i.Label()+"\t"+i.PrettyDone()+"\t"+color+"("+strconv.Itoa(pt)+")"+"\t\t"+extra+i.Text+restText)
	}

	// flsuh out the writer buffer
	w.Flush()
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
