# Go Notes
Here are all of my notes that relate to Go that I took while writing my command line application
  1. [General Go Things](#general-go-things)
  2. [Variables](#variables)
  3. [Errors](#errors)
  4. [For loops](#for-loops)
  5. [Structs](#structs)
  6. [Functions](#functions)
  7. [If Statements](#if-statements)




## General Go things
    1. [File Setup](#file-setup)
    2. [Imports](#imports)

### File Setup
- Go is all about the 'main'
- the main file is called `main.go`, that is in the main package, it's function is called `main()` and this is called after all the `init()` functions are called in all the sub files
- For each file, you want to include the
  1. package the file belongs to
  2. the imports to the file
  3. the functions and variables for the file
- the things you import (usually directories) will ahve variables with names corresponding to the last directory in the file path of the import statement

### Imports
- ___"encoding/json"___ allows us to use json to serialize data
  - __Marshal(object)__ will serialize
- ___"io/util"___ allows us to read/write to files
- ___"fmt"___ allows us to print to screen
- ___"os"___ allows us to interact with the system
  - __Exit(1)__ exits the program with a fail code
  - __WrtieFile(filename string, b []bytes, perm os.FileMode)
    - b is the thing you want to write, the stream of bytes
    - perm is the permissions: 0 (for octal), ### for user, group, and other permissions
- ___"github.com/spf13/cobra"___ allows us to use cobra's commands to set up the application



## Variables
- `:=` is used to declare and assign a variable in one go (you don't have to give it a type this way)
- It sure seems like the declaration of variables is backwards (name Type)
- Arrays are declared as []array
- If you don't use `:=`, you have to give the variable a type (or just var)
- You can do pointers with & (that makes the pointer) and dereference with *



## Errors
- These are just values, not exceptions
- Always do error handling
- You will often use `if err != nil {return err}`
  - `nil` is like null, obviously
- You usually try to do this by giving a function an error return type in addition to whatever other types you need



## For loops
- You can use the `for x,y := range ___` in order to do a for each thing
  - x is the ___index___ or ___key___, y is the __value__ for that key
    - If you don't want to give one of those a value, you can use `_`
  - then you have to put the thing you want to loop over



## Structs
- These are your own data type that you can create as you desire
- There are no constructors



## Functions
- Functions are decalred with `func`
- You do not have to give a return type to the function
- Functions tend to be capitalized 
- You can give a return type, it just goes at the end of the declaration: `func Name(name Type) returnType {}`
- Functions can have multiple returns, so you can also assign multiple variables from one function



## If statements
- You do them `if condition {}`, without any parentheses
- You can put a declaration of a variable in the if statement if you are really only going to use it right then:
```go
if err := json.Unmarshal(b, &items); err != nil {
            return []Item{}, err
                
}
// this declares the err for just that statement
// notice how it gives the function a pointer, since you don't get that info from running the function when it is in the if statement
```





