package main

import (
	"fmt"

	linkedList "github.com/mertess/GoGenericLinkedList/linkedList"
)

func main() {
	var linkedList linkedList.ILinkedList[int] = &linkedList.LinkedList[int]{}
	linkedList.AddToHead(10)
	linkedList.AddToTail(15)
	linkedList.PopTail()
	fmt.Println(*linkedList.PeekHead())
	fmt.Println(*linkedList.PeekTail())
	fmt.Println(linkedList.GetLength())
	fmt.Println(linkedList.GetValues())
}
