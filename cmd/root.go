/*
Copyright © 2023 Bryce Tolman
*/
package cmd

import (
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// file names for later
var dataFile string
var errorFile string
var cfgFile string

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

	// get the log file open
	file, err := os.OpenFile(errorFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	// find the home directory of the user
	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile and error file using --errorfile.")
	}

	// set the dataFile
	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+".todos.json", "data file to store todos")

	// set the errorFile
	rootCmd.PersistentFlags().StringVar(&errorFile, "errorfile", home+string(os.PathSeparator)+".todos_errors.log", "file to store the errors log from running the program")

	//------------- Create Files ------------------------------------

	// make the files if they don't exist
	if !fileExists(dataFile) {
		_, err := os.Create(dataFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	if !fileExists(errorFile) {
		_, err := os.Create(errorFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	// -----------------------------------------------------------------

}
