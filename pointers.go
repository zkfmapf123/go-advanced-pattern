package best_pattern

import (
	"fmt"
	"unsafe"
)

func Pointers() {

	// unsafe
	// arr := []int{1, 2, 3}
	// for _, v := range arr {
	// 	inpt := unsafe.Pointer(&v)
	// 	fmt.Println(reflect.TypeOf(v))
	// 	fmt.Println((*int)(inpt))      // Address
	// 	fmt.Println(*(*int)(inpt))     // Value
	// 	fmt.Println(*(*float64)(inpt)) // garbage _1
	// 	fmt.Println(*(*string)(inpt))  // garbage _2
	// }

	// safe
	// int32 => 4, int64 => 8
	safeArr := []int32{1, 2, 3, 4, 5, 6}
	// 0번째의 주소 값
	arrPtr := unsafe.Pointer(&safeArr[0])
	nextPtr := (*int)(unsafe.Pointer(uintptr(arrPtr) + 1*unsafe.Sizeof(safeArr[0])))

	fmt.Println("sizeof >> ", unsafe.Sizeof(safeArr[0]))
	fmt.Println(arrPtr, nextPtr)
}
