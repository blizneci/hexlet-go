// Colored text
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/hexlet-go/greeting"
)

var colors = map[string]color.Attribute{
	"белый":     color.FgWhite,
	"красный":   color.FgRed,
	"зеленый":   color.FgGreen,
	"желтый":    color.FgYellow,
	"синий":     color.FgBlue,
	"пурпурный": color.FgMagenta,
	"голубой":   color.FgCyan,
}

var backgrounds = map[string]color.Attribute{
	"черный":    color.BgBlack,
	"красный":   color.BgRed,
	"зеленый":   color.BgGreen,
	"желтый":    color.BgYellow,
	"синий":     color.BgBlue,
	"пурпурный": color.BgMagenta,
	"голубой":   color.BgCyan,
	"белый":     color.BgWhite,
}

var attributes = map[string]color.Attribute{
	"жирный":        color.Bold,
	"тусклый":       color.Faint,
	"курсив":        color.Italic,
	"подчеркивание": color.Underline,
	"мигающий":      color.BlinkSlow,
	"скрытый":       color.Concealed,
	"зачеркнутый":   color.CrossedOut,
}

var (
	errExit   = errors.New("выход")
	errNoText = errors.New("не введен текст")
)

// Text is an object to work with
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

func printMap(mapColors map[string]color.Attribute) {
	for key, val := range mapColors {
		if key == "скрытый" {
			val = color.FgWhite
		}
		c := color.New(val)
		c.Printf("%s", key)
		fmt.Print(" ")
	}
	fmt.Print("\n\n")
}

func printPossibleColors() {
	fmt.Println("Приложение может применять к тексту следующие цвета:")
	printMap(colors)
	fmt.Println("Приложение может применять к тексту следующие цвета фона:")
	printMap(backgrounds)
	fmt.Println("Приложение может применять к тексту следующие аттрибуты:")
	printMap(attributes)
	fmt.Println("Цвета текста, фона и аттрибуты можно смешивать.")
}

func processInpText(r *bufio.Reader, tPtr *Text) error {
	buf, err := r.ReadSlice('\n')
	if err != nil {
		panic(err)
	}
	tPtr.text = string(buf[:len(buf)-1])
	if tPtr.text == "" {
		return errNoText
	}
	if tPtr.text == "q" {
		return errExit
	}
	return nil
}

func processFgColor(r *bufio.Reader, tPtr *Text) error {
	buf, err := r.ReadSlice('\n')
	if err != nil {
		panic(err)
	}
	if string(buf[:len(buf)-1]) == "q" {
		return errExit
	}
	tPtr.fgColor = strings.Split(string(buf[:len(buf)-1]), " ")
	return nil
}

func processBgColor(r *bufio.Reader, tPtr *Text) error {
	buf, err := r.ReadSlice('\n')
	if err != nil {
		panic(err)
	}
	if string(buf[:len(buf)-1]) == "q" {
		return errExit
	}
	tPtr.bgColor = strings.Split(string(buf[:len(buf)-1]), " ")
	return nil
}

func processAttr(r *bufio.Reader, tPtr *Text) error {
	buf, err := r.ReadSlice('\n')
	if err != nil {
		panic(err)
	}
	if string(buf[:len(buf)-1]) == "q" {
		return errExit
	}
	tPtr.attr = strings.Split(string(buf[:len(buf)-1]), " ")
	return nil
}

func printColoredText(tPtr *Text) {
	coloredText := color.New()
	for _, fgCol := range tPtr.fgColor {
		if val, ok := colors[fgCol]; ok {
			coloredText.Add(val)
		}
	}
	for _, bgCol := range tPtr.bgColor {
		if val, ok := backgrounds[bgCol]; ok {
			coloredText.Add(val)
		}
	}
	for _, attr := range tPtr.attr {
		if val, ok := attributes[attr]; ok {
			coloredText.Add(val)
		}
	}
	coloredText.Print(tPtr.text)
	fmt.Println()
}

func main() {
	greeting.SayHello()
	greeting.Introduce()
	printPossibleColors()
mainCycle:
	for {
		r := bufio.NewReader(os.Stdin)
		t := Text{}
		fmt.Println("Введите текст (q для выхода из приложения):")
		var err error
		err = processInpText(r, &t)
		if err != nil {
			switch err {
			case errNoText:
				fmt.Println(err.Error())
				continue
			case errExit:
				exit()
				break mainCycle
			}
		}
		fmt.Println("Введите цвет текста (q для выхода из приложения):")
		err = processFgColor(r, &t)
		if err == errExit {
			exit()
			break mainCycle
		}
		fmt.Println("Введите цвет фона текста (q для выхода из приложения):")
		err = processBgColor(r, &t)
		if err == errExit {
			exit()
			break mainCycle
		}
		fmt.Println("Введите аттрибуты текста (q для выхода из приложения):")
		err = processAttr(r, &t)
		if err == errExit {
			exit()
			break mainCycle
		}

		fmt.Printf("Введенный текст: %s\n", t.text)
		fmt.Printf("Введенный цвет текста: %s\n", t.fgColor)
		fmt.Printf("Введенные цвет фона: %s\n", t.bgColor)
		fmt.Printf("Введеный аттрибут: %s\n", t.attr)

		printColoredText(&t)

	}
}

//
