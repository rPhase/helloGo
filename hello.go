package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func toInt(str string) int {
	fmt.Printf("Converting string to int...\n")
	tmp, err := strconv.Atoi(str)
	if err == nil {
		fmt.Println(tmp)
		return tmp
	}
	fmt.Fprintln(os.Stderr, "parsing int:", err)
	return 0
}

func read(scanner *bufio.Scanner) string {
	fmt.Printf("Reading...\n")
	// scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		return line
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return ""
}

func main() {
	fmt.Printf("こんにちは世界。\n")
	scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println(read(scanner))
	n := toInt(read(scanner))
	fmt.Println(n*n)
}
