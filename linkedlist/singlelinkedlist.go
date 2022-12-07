package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head *node
	tail *node
}

func (l *linkedlist) isEmpty() {
	if l.head == nil {
		fmt.Println("is Empty")
	}
}

func (l *linkedlist) push(n int) {
	newnode := node{n, nil}
	if l.head == nil {
		l.head = &newnode
	} else {
		l.tail.next = &newnode
	}
	l.tail = &newnode
}

func (l *linkedlist) display() {
	p := l.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next
	}
}

func (l *linkedlist) len() int {
	p := l.head
	count := 0
	for p != nil {
		p = p.next
		count++
	}
	return count

}

func (l *linkedlist) addfirst(n int) {
	newnode := &node{n, nil}
	if l.head == nil {
		l.head = newnode
		l.tail = newnode
	} else {
		newnode.next = l.head
		l.head = newnode
	}

}

func (l *linkedlist) addlast(n int) {
	newnode := &node{n, nil}
	if l.head == nil {
		l.head = newnode
		l.tail = newnode

	} else {
		l.tail.next = newnode
	}
}

func (l *linkedlist) addany(n int, pos int) {
	newnode := &node{n, nil}
	p := l.head
	i := 0
	for i < pos-1 {
		p = p.next
		i++
	}
	newnode.next = p.next
	p.next = newnode
}

func main() {
	l := linkedlist{}
	l.isEmpty()
	fmt.Println()
	l.push(10)
	l.push(20)
	l.push(30)
	l.push(40)
	l.display()
	fmt.Print()
	k := l.len()
	fmt.Print()
	fmt.Println()
	fmt.Print(k)
	fmt.Println()
	l.addfirst(5)
	l.addfirst(6)
	l.display()
	fmt.Println()
	l.addlast(9)
	l.display()
	fmt.Println()
	l.addany(0, 4)
	l.display()

}
