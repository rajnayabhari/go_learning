// Packages in Go are like libraries or modules in other programming languages.
// You can package your code and reuse it somewhere else.
// The source code of a package can be distributed in more than one .go file.
// So far, we've been writing the main package and have made a few references to other native packages.
// In this section, you'll learn what a package is. You'll also learn how to create one
//  and how to consume external packages.


// Main package import importance
// As you might have noticed, even the most straightforward program in Go has to be part of a package. 
// Usually, the default package is the main package, the one we've been using so far. 
// If a program is part of the main package, Go generates a binary file. 
// When that file runs, it calls the main() function.
