package _5_array

import (
	"fmt"
	"testing"
)

func TestNewArray(t *testing.T) {
	array := NewArray(5)
	array.Print()
	fmt.Println("len of array: ", array.Len())
	fmt.Println(array.isIndexOutOfRange(0))

}
