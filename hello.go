/*
*Author AMRiezz
*/

/*
*This is comment
*/

/* my first program in Go */
package main

/*The first line of the program package main defines the package name in which this program should lie. It is a mandatory statement,as Go programs run in packages. The main package is the startingpoint to run the program. Each package has a path and name associated with it.*/

import "fmt"
/*The next line import "fmt" is a preprocessor command which tells the Go compiler to include the files lying in the package fmt.*/

func main() {
   /* This is my first sample program. */
   fmt.Println("Hello, World!")
}

/*The next line fmt.Println(...) is another function available in Go which causes the message "Hello, World!" to be displayed on the screen. Here fmt package has exported Println method which is used to display the message on the screen.*/

/*go run hello.go*/

