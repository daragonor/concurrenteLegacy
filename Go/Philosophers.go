package main

import (
	"fmt"
	"time"
)

var end chan bool

func philosopher(id int, fi, fi1 chan bool) {
	for {
		fmt.Printf("think %d\n", id)
		time.Sleep(time.Millisecond * 500)
		<-fi
		<-fi1
		fmt.Printf("eat %d\n", id)
		time.Sleep(time.Millisecond * 500)
		fi <- true
		fi1 <- true
		//close(forks)
	}
	end <- true
}

func fork(fi chan bool) {
	for {
		fi <- true
		time.Sleep(time.Millisecond * 500)
		<-fi
	}
}

func main() {
	n := 5
	forks := make([]chan bool, n)
	for i := range forks {
		forks[i] = make(chan bool, 1)
	}
	i := 0
	for i < 5 {
		go philosopher(i, forks[i], forks[(i+1)%n])
		go fork(forks[i])
		i++
	}
	for i := 0; i < n; i++ {
		<-end
	}
}
