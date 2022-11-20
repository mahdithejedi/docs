package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

// __source__ = 'https://go.dev/blog/pipelines'

func printer(value string, num int) {
	fmt.Println("!", value, "!, runner", num)
}

type runnerFuncTemp func(string, int)

func initRunners(counter int, runFunc runnerFuncTemp) []runnerFuncTemp {
	runners := make([]runnerFuncTemp, counter)
	for i := 0; i < counter; i++ {
		runners[i] = runFunc
	}
	return runners
}

func runner(counter int, runFunc runnerFuncTemp, fin <-chan bool, data <-chan string) {
	runnerCounter := 0
	runnerArray := initRunners(counter, runFunc)
	for true {
		select {
		case _data := <-data:
			runnerArray[runnerCounter](_data, runnerCounter)
			runnerCounter += 1
			if runnerCounter > counter-1 {
				runnerCounter = 0
			}
		case <-fin:
			return
		}
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	dataStream := make(chan string)
	finish := make(chan bool, 1)
	go func() {
		for true {
			select {
			case _, ok := <-finish:
				if ok {
					break
				}
			default:
				dataStream <- RandStringRunes(rand.Intn(256))
			}
		}
	}()
	go runner(20, printer, finish, dataStream)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	//go func() {
	for _ = range c {
		finish <- true
		close(dataStream)
		close(finish)
		fmt.Println("out!!!!")
		return
	}
	//}()
}
