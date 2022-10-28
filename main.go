package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}

	qntt := count(src)

	fmt.Println(qntt)
}

func count(src string) int {

	text := strings.Split(src, " ")

	return len(text)
}

func readInput() (src string, err error) {

	flag.Parse()

	src = strings.Join(flag.Args(), "")
	if src == "" {
		return src, errors.New("missing string to match")
	}
	return src, nil
}

func fail(err error) {
	fmt.Println("count:", err)
	os.Exit(1)
}
