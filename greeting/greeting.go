// Package greeting implements a Hello function to return string to the main.
package greeting

import "fmt"

const greeting = "Golang for Brave!"

// Hello function returns string to the main function
func SayHello() {
	fmt.Println(greeting)
}

// Prints description
func Introduce() {
	fmt.Println("Это приложение раскрашивает введенный текст в указанные цвета")
	fmt.Println("Поддерживается изменение цвета текста, цвета фота текста и дополнительные аттрибуты.")
}

//
