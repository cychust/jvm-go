package rtda

import "jvm/ch06/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}