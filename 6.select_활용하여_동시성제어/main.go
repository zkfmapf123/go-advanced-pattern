package main

import (
	"fmt"
)

type Routine struct {
	msg    string
	isWait bool
}

type RoutineOpts func(rou *Routine)

func setMsg(msg string) RoutineOpts {
	return func(rou *Routine) {
		rou.msg = msg
	}
}

func setWait(rou *Routine) {
	rou.isWait = true
}

func deafultRoutine() *Routine {
	return &Routine{
		msg:    "",
		isWait: false,
	}
}

func NewRoutine(rouOpts ...RoutineOpts) *Routine {
	rou := deafultRoutine()

	for _, r := range rouOpts {
		r(rou)
	}
	return rou
}

// ////////////////////////////////////////////////////////////////////////// Routine ////////////////////////////////////////////////////////////////////////////
func setRoutine(msg string, num int) chan Routine {
	c := make(chan Routine)

	go func() {
		for i := 0; i < num; i++ {

			// 짝수는 Wait is true
			if i%2 == 0 {
				c <- *NewRoutine(setMsg(fmt.Sprintf("%s -> %d", msg, i)), setWait)
			} else {
				c <- *NewRoutine(setMsg(fmt.Sprintf("%s -> %d", msg, i)))
			}
		}
	}()

	return c
}

func fanin(rousCh ...chan Routine) chan Routine {
	c := make(chan Routine)

	for _, r := range rousCh {

		go func(rr chan Routine) {
			for {
				c <- <-rr
			}
		}(r)
	}

	return c
}

func main() {
	rous1, rous2 := setRoutine("hello", 5), setRoutine("bye", 5)
	f := fanin(rous1, rous2)

	trueCount, falseCount := 0, 0

	for i := 0; i < 5; i++ {
		select {
		case s := <-f:
			{
				if s.isWait {
					trueCount += 1
				} else {
					falseCount += 1
				}
			}
		}
	}

	fmt.Printf("result true : %d, false : %d\n", trueCount, falseCount)
}
