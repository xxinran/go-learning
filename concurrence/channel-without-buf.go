package main

import (
	"bufio"
	"fmt"
	"os"
)

func printInput(ch chan string) {
	for val := range ch {
		if val == "EOF" {
			break
		}
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan string)
	defer close(ch)
	go printInput(ch)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		ch <- val
		if val == "EOF" {
			break
		}
	}
}
