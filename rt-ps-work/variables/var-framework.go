package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var ( //functions declared outside of main are "Global"
	module = "4" //Needs to be an integer
	clip   = 2   //needs to be an integer
)

func main() {

	// courseComplete := false
	name := os.Getenv("USER")
	course := "Getting Started with Go"

	fmt.Println("Name and course are set to", name, "and ", course, ".")
	fmt.Println("Module and clip are set to", module, "and ", clip, ".")
	fmt.Println("Name is of type ", reflect.TypeOf(name))
	fmt.Println("Module is of type ", reflect.TypeOf(module))
	// total := module + clip
	// fmt.Println("Module plus clip equals", total)

	// variables declared in the function are local only to the function
	iModule, err := strconv.Atoi(module) //Go passes a "copy" of the variable value to the 2nd function
	if err == nil {
		total := iModule + clip
		fmt.Println("Module plus clip equals", total)
	}

	// fmt.Println("Memory address of *course* variable is", &course)

	// var ptr *string = &course
	// fmt.Println("Pointing course variable at address", ptr, "which holds this value:", *ptr, ".... which matches the original:", course)

	fmt.Println("\nHi", name, "your current course is", course)
	updateCourseByVal(course)

	fmt.Println("You're currently watching", course)

	updateCourseByRef(&course)

	fmt.Println("You're currently watching", course)

}

func updateCourseByVal(course string) string {

	course = "Getting Started with Docker"
	fmt.Println("Updated course to ", course)
	return course

}

func updateCourseByRef(course *string) {

	*course = "Getting Started with Docker"
	fmt.Println("Updated course to ", *course)
	return //*course

}
