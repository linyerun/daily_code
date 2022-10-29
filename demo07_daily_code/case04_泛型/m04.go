package main

import (
	"fmt"
)

//做一个双向链表

type linkedList[E any] struct {
	head, tail *elem[E]
}

func NewLinkedList[E any]() *linkedList[E] {
	list := &linkedList[E]{
		head: new(elem[E]),
		tail: new(elem[E]),
	}
	list.head.next, list.tail.prev = list.tail, list.head
	return list
}

func (t *linkedList[E]) PushFirst(e E) {
	elem := newElem(e)
	head := t.head
	head.next, head.next.prev, elem.next, elem.prev = elem, elem, head.next, head
}

func (t *linkedList[E]) PushLast(e E) {
	elem := newElem(e)
	tail := t.tail
	tail.prev, tail.prev.next, elem.next, elem.prev = elem, elem, tail, tail.prev
}

func (t *linkedList[E]) PollFirst() (e E, err error) {
	if t.head.next == t.tail {
		err = fmt.Errorf("list's size is zero")
		return
	}
	h, p := t.head, t.head.next
	e = p.value
	h.next, p.next.prev = p.next, h //移除元素
	return
}

func (t *linkedList[E]) PollLast() (e E, err error) {
	if t.head.next == t.tail {
		err = fmt.Errorf("list's size is zero")
		return
	}
	tail, p := t.tail, t.tail.prev
	e = p.value
	tail.prev, p.prev.next = p.prev, tail //移除元素
	return
}

func (t *linkedList[E]) Size() (size int) {
	p := t.head
	for p.next != t.tail {
		size++
		p = p.next
	}
	return
}

func (t *linkedList[E]) removeElem(e *elem[E]) {
	e.next, e.prev = e.prev.next, e.next.prev
}

type elem[E any] struct {
	value      E
	prev, next *elem[E]
}

func newElem[E any](e E) *elem[E] {
	return &elem[E]{
		value: e,
	}
}

//使用这个双向链表做一个LRUCache

type KV[K comparable, V any] struct {
	key   K
	value V
}

func NewKV[K comparable, V any](key K, val V) *KV[K, V] {
	return &KV[K, V]{
		key: key,
		val: val,
	}
}

type LRUCache[K comparable, V any] struct {
	list *linkedList[*KV[K, V]]
	m    map[K]*elem[*KV[K, V]]
	cap  int
}

func NewLURCache[K comparable, V any](cap int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		list: NewLinkedList[*KV[K, V]](),
		m:    make(map[K]*elem[*KV[K, V]]),
		cap:  cap,
	}
}

func (t *LRUCache[K, V]) Get(key K) (val V, ok bool) {
	elem, ok := t.m[key]
	if !ok {
		return
	}
	val = elem.value.value
	//将elem移动到顶端
	t.list.removeElem(elem)      //除
	t.list.PushFirst(elem.value) //添
	return
}

func (t *LRUCache[K, V]) Push(key K, value V) {
	if elem, ok := t.m[key]; ok {
		//该点存在
		t.list.removeElem(elem)             //除
		t.list.PushFirst(NewKV(key, value)) //加
		return
	}
	if len(t.m) == t.cap {
		//移除一个元素
		if elem, err := t.list.PollLast(); err != nil {
			panic(err)
		} else {
			delete(t.m, elem.key)
		}
	}
	//点不存在
	elem := newElem(NewKV(key, value))
	t.m[key] = elem
	t.list.PushFirst(elem.value)
}

func main() {
	list := NewLinkedList[int]()
	list.PushFirst(10)
	list.PushFirst(20)
	list.PushLast(10)
	list.PushLast(20)
	size := list.Size()
	println("list-size:", size)
	for i := 0; i < size; i++ {
		e, err := list.PollFirst()
		if err != nil {
			panic(err)
		}
		fmt.Print(e, " ")
	}
	println()
	list.PushFirst(1)
	list.PushFirst(2)
	list.PushLast(3)
	list.PushLast(4)
	size = list.Size()
	println("list-size:", size)
	for i := 0; i < size; i++ {
		e, err := list.PollLast()
		if err != nil {
			panic(err)
		}
		fmt.Print(e, " ")
	}
	println()
	size = list.Size()
	println("list-size:", size)
}
