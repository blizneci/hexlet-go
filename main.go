// Colored text
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/hexlet-go/greeting"
)

const instruction = `Это приложение для вывода текста, раскрашенного в разные цвета.
Можно указать цвет текста, фона и задать дополнительные аттрибуты.

`

colors := map[string]color.Attirbute{
	"белый": color.FgWhite,
	"красный": color.FgRed,
	"зеленый": color.FgGreen,
	"желтый": color.FgYellow,
	"синий": color.FgBlue,
	"пурпурный": color.FgMagenta,
	"голубой": color.FgCyan,
}

backgrounds := map[string]color.Attirbute{
	"черный": color.BgBlack,
	"красный": color.BgRed,
	"зеленый": color.BgGreen,
	"желтый": color.BgYellow,
	"синий": color.BgBlue,
	"пурпурный": color.BgMagenta,
	"голубой": color.BgCyan,
	"белый": color.BgWhite,
}

attributes := map[string]color.Attribute{
	"жирный": color.Bold,
	"тусклый": color.Faint,
	"курсив": color.Italic,
	"нижнее подчеркивание": color.Underline,
	"мигающий медленно": color.BlinkSlow,
	"мигающий быстро": color.BlinkRapid,
	"обратное видео": color.ReverseVideo,
	"скрытый": color.Concealed,
	"зачеркнутый": color.CrossedOut,
}

type Text struct {
	text    string
	fgColor []string
	bgColor []string
	attr    []string
}

func exit() {
	fmt.Println("Работа приложения завершается.")
	time.Sleep(1 * time.Second)
}

func inputText(r *bufio.Reader) string {
	var buf []byte
	buf, _ = r.ReadSlice('\n')
	return string(buf[:len(buf)-1])
}
func getText(r *Reader) Text, error {
	t := Text{}
	fmt.Println("Введите текст:")
	t.text = inputText(r)
	if t.text == "q" {
		err := errors.New("")
		return Text{}, 
	}
	
}

func main() {
	greeting.SayHello()
	greeting.Introduce()
	r := bufio.NewReader(os.Stdin)

	fmt.Printf("Введенный текст: %s\n", t.text)
	fmt.Printf("Введенный цвет текста: %s\n", t.fgColor)
	fmt.Printf("Введенные цвет фона: %s\n", t.bgColor)
	fmt.Printf("Введеный аттрибут: %s\n", t.attr)


	/*
		fmt.Println("Text with colors")
		for _, col := range colors {
			c := color.New(col)
			c.Printf("text\n")
		}

		fmt.Println("Text with background")
		for _, bg := range backgrounds {
			b := color.New(bg)
			b.Print("text")
			fmt.Println()
		}

		fmt.Println("Text with attributes")
		for _, attr := range attributes {
			a := color.New(attr)
			a.Print("text")
			fmt.Println()
		}
		color.New().Add(color.Underline).Println("Bold text")
	*/

}

//
