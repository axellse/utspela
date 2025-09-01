package main

import (
	"fmt"
	"os"
)

var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"
var Reset = "\033[0m"

func PrintColor(text string, color string) {
	fmt.Println(color + text + Reset)
}

func Done(text string) {
	fmt.Print(text + " [")
	fmt.Print(Green + "DONE" + Reset)
	fmt.Println("]")
}

func Info(text string) {
	fmt.Print(text + " [")
	fmt.Print(Blue + "INFO" + Reset)
	fmt.Println("]")
}

func FError(text string) {
	Error(text)
	os.Exit(1)
}

func Error(text string) {
	fmt.Print(text + " [")
	fmt.Print(Red + "ERROR" + Reset)
	fmt.Println("]")
}