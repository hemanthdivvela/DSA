package main

import "fmt"

type node struct {
	data int
	next *node
}

type stacklist struct {
	head *node
}

func (s *stacklist) push(n int) {
	newnode := &node{n, nil}
	if s.head == nil {
		s.head = newnode
	} else {
		newnode.next = s.head
		s.head = newnode
	}
}

func (s *stacklist) display() {
	p := s.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next
	}
}

func (s *stacklist) pop() {
	if s.head == nil {
		fmt.Print("it is a empty stack")
		return
	}
	// s.head.data = s.head.data
	s.head = s.head.next
}

func (s *stacklist) removedelement() {
	if s.head == nil {
		fmt.Print("it is a empty stack")
		return
	}
	fmt.Print(s.head.data)
}

func (s *stacklist) len() int {
	p := s.head
	count := 0
	for p != nil {
		p = p.next
		count++
	}
	return count
}

func main() {
	l := stacklist{}
	l.push(10)
	l.push(20)
	l.push(30)
	l.display()
	fmt.Println()

	l.pop()
	l.display()
	fmt.Println()

	l.push(30)
	l.push(40)
	l.display()
	fmt.Println()

	l.removedelement()

	a := l.len()
	fmt.Println()
	fmt.Print(a)

}
