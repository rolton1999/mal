package main

import (
	"bufio"
	"fmt"
	"mal/src/printer"
	"mal/src/reader"
	"os"
)

func read(s string) string {
	return reader.ReadStr(s)
}

func eval(s string) string {
	return s
}

func print(s string) string {
	return printer.Print(s)
}

func rep(s string) string {
	return print(eval(read(s)))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("user> ")
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(rep(s))
		fmt.Print("user> ")
	}
}
