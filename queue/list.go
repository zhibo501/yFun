package queue

type Value struct {
	Key   string
	Value interface{}
}

type List interface {
	Push(k string, v interface{})
	Pop() *Value
	Count() int32
	Show()
}

type lNode struct {
	prev *lNode
	next *lNode
	data *Value
}
