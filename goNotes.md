# Go Notes
Here are all of my notes that relate to Go that I took while writing my command line application

## General Go things
- Go is all about the 'main'
- the main file is called `main.go`, that is in the main package, it's function is called `main()` and this is called after all the `init()` functions are called in all the sub files
- For each file, you want to include the
  1. package the file belongs to
  2. the imports to the file
  3. the functions and variables for the file
- the things you import (usually directories) will ahve variables with names corresponding to the last directory in the file path of the import statement

## Variables
- `:=` is used to declare and assign a variable in one go

## Errors
- 

## For loops
- You can use the `for x,y := range ___` in order to do a for each thing
  - x is the ___index___ or ___key___, y is the __value__ for that key
    - If you don't want to give one of those a value, you can use `_`
  - then you have to put the thing you want to loop over

