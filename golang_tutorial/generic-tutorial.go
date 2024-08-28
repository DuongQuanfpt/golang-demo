package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	next *List[T]
	val  T
}

func(xs *List[T]) String() string {
	if xs.next == nil {
		return fmt.Sprintf("{ val: %v, next: <nil> }", xs.val)
	} else {
		return fmt.Sprintf("{ val: %v, next: %s }", xs.val, fmt.Sprint(xs.next))
	}
}

func (xs *List[T]) append (val T) *List[T] {
	
	if(xs.next == nil){
		xs.next = NewList(val)
	} else {
		ys := xs
		ys.next.append(val)
	}
	return xs
}

func NewList[T any](val T) *List[T] {
	xs := new(List[T])
	xs.val = val
	xs.next = nil
	return xs
}


func genericTutorial() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))//2

	ss := []string{"hello", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))//0

	xs := NewList(0)
	fmt.Println(xs)
	
	for i := 1; i < 5; i++ {
		xs.append(i)
		fmt.Println(xs)
	}
}

