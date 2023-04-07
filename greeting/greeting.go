// Package greeting implements a Hello function to return string to the main.
package greeting

import "fmt"

const (
	greeting    = "Golang for Brave!"
	description = `Это приложение раскрашивает введенный текст в указанные цвета.
Поддерживается изменение цвета текста, цвета фона текста и дополнительные аттрибуты.`
)

// Hello function returns string to the main function
func SayHello() {
	fmt.Println(greeting)
}

// Prints description
func Introduce() {
	fmt.Println(description)
}

//
