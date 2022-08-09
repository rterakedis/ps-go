# Pluralsight Notes from Go Fundamentals

---

## Go History

Mascot:  the Go Gopher (by Renee French)
Created by:   Robert Griesemer, Ken Thompson, Rob Pike
History
  - Started in ~2007
  - First Release in 2009 (Open Sourced BSD Licensed)
  - Notable Projects:   Docker, Kubernetes, etcd, Terraform
  - Designed as Systems Language --> OS and Infrastructure
  - Evolved as all-arounder for Web Services, etc.
  - Similar to C, but simpler -- More concise and does it's own garbage collection
  - Only 25 keywords
  - Cross Platform -- Compiles programs for multiple OS's, and can be installed on multiple OS's
  - Compiled Language --> runs without an interpreter (Fast!!)
  - Strongly Typed (particularly variables)
  - Supports Type Inference





---

## Installing Go

Mac, Windows, Linux native installers.  Anything else install from Source.

### VS Code Configuration

- 

### Go Version Check

`go version`





---

## Go Structure

```
Module (dependency & version tracking)
|-Package (folder)
    |-source-code1.go
    |-source-code2.go
```

### Initialize Go Module

From the folder where you intend to start building your module, Initialize the Module (Using the name of the URL where you'll host your code)

From within `/Users/rterakedis/Documents/Git-Repos/ps-go/rt-ps-work`, run `go mod init github.com/rterakedis/ps-go/rt-ps-work`





---

## Starting a New Go Program

Every go program needs a package declaration.

Every file that starts with `package main` -- this tells Go to compile this code as an application (something clickable or standalone) and not a shared library.





---

## Functions

- Go passes arguments by value, NOT by reference

```
func <name> (params) (returns) {

    <code>

}
```

The entry point is `func main()` -- takes no arguments and returns no values (the entire program exits with 0 or non-zero [error] code).





---

## Variables

**Variable Rules:**
- Must start with a letter
- can't be go keyword
- can't contain spaces
- can't contain special characters
- Variables Declared at the Package Level are allowed to NOT be used
- Variables declared within a function MUST be used or Go throws an error
- `var <name> <type>`

**Best Practices**
- use meaningful names
- keep names short
- use camelCase
- the `&` sign returns the memory address of the variable -- i.e. `&course` --> the memory address of the `course` variable.
- Pointers let you pass by reference (rather than by value)

### Short Assignment

Use the `:=` characters to short-assign a variable with type inference
> Short Assignment is ONLY useable inside a function, and not for global variables.

### Examples 
```
var (
    name, course    string
    module, clip    int
)
```

```
var (
	name 	string
	course 	string
	module 	int
	clip 	int
)
```

### Type Inference 

```
var (
    name = "Nigel Poulton"
    course = "Getting Started with Go"
    module = 4  //Needs to be an integer
    clip = 2  //needs to be an integer
)
```

```
var (
	name, course = "Nigel Poulton", "Getting Started with Go"
	module, clip = 4, 2
)
```

```
	// Short Assignment with type inference
    name := "Nigel Poulton"
	course := "Getting Started with Go"
```



---

## Pointer Variables

Pointers on Pointer Variables:

- The `*` indicates a pointer variable --> references the memory of another variable
- The `&` indicates you want the memory address of the variable and not the value of the variable 

`var ptr *string = &course`

In other words, the `ptr` variable is a string pointing to the memory address of the `course` variable.

To de-reference the `ptr` variable (i.e. get the value stored at the reference), use the variable with the asterisk --> `*ptr`

```
var ptr *string = &course
	fmt.Println("Pointing course variable at address", ptr, "which holds this value:", *ptr, ".... which matches the original:", course)
``` 
yields the following output:   `Pointing course variable at address 0x112f0b0 which holds this value: Getting Started with Go .... which matches the original: Getting Started with Go`


### Passing Variables by Value

```
func main {
	name := "Nigel Poulton"
	course := "Getting Started with Go"
    
    fmt.Println("\nHi", name, "your current course is", course)
    updateCourseByVal(course)
    fmt.Println("You're currently watching", course)
}
func updateCourseByVal(course string) {

	course = "Getting Started with Docker"
	fmt.Println("Updated course to ", course)
	return

}
```
Yields the following: 
```
Hi Nigel Poulton your current course is Getting Started with Go
Updated course to  Getting Started with Docker
You're currently watching Getting Started with Go
```

> **NOTE:** The change to the "course" variable is localized to the `updateCourseByValue` function and therefore the main function's variable value is not modified.

### Passing Variables by Reference

```
func main {
	name := "Nigel Poulton"
	course := "Getting Started with Go"
    
    fmt.Println("\nHi", name, "your current course is", course)
    updateCourseByRef(&course)
	fmt.Println("You're currently watching", course)
}
func updateCourseByRef(course *string) {

	*course = "Getting Started with Docker"
	fmt.Println("Updated course to ", *course)
	return 

}
```
Yields the following: 
```
Hi Nigel Poulton your current course is Getting Started with Go
Updated course to  Getting Started with Docker
You're currently watching Getting Started with Docker
```

> **NOTE:** The `main` function calls the `updateCourseByRef` function using memory address of the `course` variable in `main`.   The `update...` function stores that memory address in the pointer variable called `course` scoped to that `update...` function.   The code de-references the `*course` variable to assign the new string at that memory location, and then prints the dereferenced variable.





---

## Constants

Constants **CAN NOT** be changed --> "Immutable variables"

`const c = 186000`



---

## Special Go Keywords


| word | word | word | word | word |
| ---- | ---- | ---- | ---- | ---- |
| break | default | func | interface | select |
| case | defer | go | map | struct |
| chan | else | goto | package | switch |
| const | fallthrough | if | range | type |
| continue | for | import | return | var |




---

## Importing Functions

Extends functionality by adding functions from additional packages:

```
import (
	"fmt" //text formatting
	"reflect"  //reflection to determine type of
	"strconv"  //string conversion to Integers
    "os" //OS & Environment Variables
)
```





---

## Compiling & Testing

- Testing for dev:  `go run <filename>.go`
- Testing Compiled App (for shipping to customers):   `go build <packagename>` or `go build` from within the package folder
- Install Locally Compiled App:  `go install` from within the package folder

> Change the default Bin/Path locations using `go env -w GOPATH="<path>"` and `go env -w GOBIN="<path>"`
