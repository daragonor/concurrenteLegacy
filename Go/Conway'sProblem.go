package main

import (
	"fmt"
)

func compress() {
	n := 0
	previous := <-inC
	for c := range inC {
		if c == previous && n < MAX-1 {
			n = n + 1
		} else {
			if n > 0 {
				pipe <- rune(n + 1 + 48)
				n = 0
			}
			pipe <- previous
			previous = c
		}
	}
	if n > 0 {
		pipe <- rune(n + 1 + 48)
	}
	pipe <- previous
	close(pipe)
}
func output() {
	m := 0
	for c := range pipe {
		//c = <-pipe
		outC <- c
		m = m + 1
		if m >= K {
			outC <- '\n'
			m = 0
		}
	}
	close(outC)
}

const MAX = 9
const K = 4

var inC, pipe, outC chan rune

func main() {
	inC = make(chan rune)
	pipe = make(chan rune)
	outC = make(chan rune)
	go compress()
	go output()
	go func() {
		msg := "AABCCCCC"
		for _, c := range msg {
			inC <- c
		}
		close(inC)
	}()
	var str string
	for c := range outC {
		str += string(c)
	}
	fmt.Println(str)
	fmt.Println()
}
