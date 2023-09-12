package best_pattern

import (
	"fmt"
	"testing"
)

type Result struct {
	data int
	err string
}

func _reliableMain() {
	input := []int{1,2,3,4}

	inputCh := _reliableGenerator(input)
	resultCh := make(chan Result ,len(input))

	go _reliableConsumber(inputCh, resultCh)

	for res := range resultCh {
		fmt.Println(res)
	}

	
	

}

func _reliableConsumber(inputCh chan int, resultCh chan Result) {
	defer close(resultCh)

	for data := range inputCh {

		if data == 2 {
			resultCh <- Result{
				data: data,
				err : "2 is invalid number",
			}
			continue
		}
		
		resultCh <- Result{
			data :data,
			err : "",
		}
	}
}

func _reliableGenerator(input []int) chan int{
	
	ch := make(chan int)

	go func() {
		defer close(ch) // 익명함수에 위치
		for _, v := range input {
			ch <- v
		}
	}()

	return ch
}

func TestReliableConsumerPattern(t *testing.T) {	
	fmt.Println("-----------------------------Reliable Consumer Pattern--------------------------------")
	_reliableMain()
}

