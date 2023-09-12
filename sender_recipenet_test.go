package best_pattern

import (
	"fmt"
	"testing"
)

// 채널을 열었다면 꼭 다시 닫아야 한다.
func sender(ch chan int) {
	for i:=0; i<10; i++ {
		ch <- i
	}

	close(ch)
	// ch <- 200 닫힌 채널의 대해서는 다시 데이터를 넣을 수 없음
}

func recipient(ch chan int) {
	for data := range ch {
		fmt.Println(data)
	}
}

func TestSenderRecipent(t *testing.T) {
	fmt.Println("---------------------------- Close Channel ----------------------------")
	
	ch := make(chan int)
	go sender(ch)
	recipient(ch)
}