package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/hexlet-go/greeting"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(greeting.Hello())
	logrus.Println("Hello, Hexlet!")
	color.Cyan("Prints text in cyan.")
	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")

	// Or just add them to New()
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("This prints bold cyan %s\n", "too!.")

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)
	err := errors.New("some critical error")
	red.Println("Warning")
	red.Printf("Error: %s\n", err)

	boldRed := red.Add(color.Bold)
	boldRed.Println("This will print text in bold red.")

	whiteBackground := red.Add(color.BgWhite)
	whiteBackground.Println("Red text with white background.")

    color.Set(color.FgYellow)

    fmt.Println(greeting.Hello())
    logrus.Println("Hello, Hexlet!")

    color.Unset()

}

//
