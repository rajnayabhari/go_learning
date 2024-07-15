package main

// The package main statement is how we tell Go that the app
// we're creating is an executable program (a file you can run).
// Our "Hello World!" app is part of the main package. A package is a set of common source code files.
//  Every executable app has this first line, even if the project or file has a different name.
// import "fmt"
// You need this import statement because you're using a function from this package to print
//
//	a message to the screen later in the program.
//
// You can include as many import statements as you want or need in your program. However, Go is idiomatic in
// this regard. If you import a package, but don't use a corresponding function from the package, the app
// won't compile.
// A great feature of Visual Studio Code is that it automatically removes unused imports in a program
// when you save the file.
// func main() {
// 	fmt.Println("Hello World!")
// }
// The func statement is a reserved word that's used to declare a function.
// This first function is named "main" because it's the starting point of our program.
// You can have only one main() function across the package main (the one you defined in the first line).
// In the main() function, you called the Println function from the fmt package.
// You sent a text message that you wanted to see on the screen.



// to run the code got to the directory where the file is and the run command go run filename.go
// and to build an executable run command go build filename.go
