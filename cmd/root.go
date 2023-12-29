/*
Copyright © 2023 Bryce Tolman
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "ToDo is a ToDo list manager application",
	Long:  `ToDo will help you get more done in less time. It is a todo list manager application designed to be very simple to use and to help you accomplish your goals.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// function to see if a file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	// make the file names
	filename1 := "C:/Users/bat20/.todos.json"
	filename2 := "C:/Users/bat20/.todos_errors.log"

	// make the files if they don't exist
	if !fileExists(filename1) {
		_, err := os.Create(filename1)
		if err != nil {
			log.Fatal(err)
		}
	}
	if !fileExists(filename2) {
		_, err := os.Create(filename2)
		if err != nil {
			log.Fatal(err)
		}
	}

	// get the log file open
	file, err := os.OpenFile(filename2, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if there is an error
	if err != nil {
		log.Fatal(err)
	}

	// keep the file from closing for a second
	defer file.Close()
	// set the log output as that file
	log.SetOutput(file)

	// Execute order 66
	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
