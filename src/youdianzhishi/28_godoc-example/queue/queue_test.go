package queue

import "fmt"

func ExampleQueue_Push() {
	q:= Queue{}
	q.Push(1,2,3,5)
	fmt.Println(q)

	// output:
	// [1 2 3 5]
}

func ExampleQueue_Pop() {
	q:= Queue{1,2,3}
	q.Pop()
	fmt.Println(q)

   // output:
   // [2 3]
}
