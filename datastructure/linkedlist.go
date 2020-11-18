// Linked List in Golang
package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	list := &Node{ // address of the new node
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.key)
		list = list.prev
	}
	fmt.Println()
}

// methrd of displaying the elements in the reverse order
/*func (l *List) Reverse() {
    curr := l.head
    var prev *Node
    l.tail = l.head

    for curr != nil {
        next := curr.next
        curr.next = prev
        prev = curr
        curr = next
    }
    l.head = prev
    Display(l.head)
}
*/
func main() {
	link := List{}
	link.Insert(5)
	link.Insert(10)
	link.Insert(15)
	link.Insert(20)
	link.Insert(25)
	link.Insert(30)

	fmt.Println("\n==============================\n")
	fmt.Printf("Head elements: %v\n", link.head.key)
	fmt.Printf("Tail elements: %v\n", link.tail.key)
	link.Display()
	// to display the elements in  reverse order in the linked list
	/*fmt.Println("\n==============================\n")
	  fmt.Printf("head: %v\n", link.head.key)
	  fmt.Printf("tail: %v\n", link.tail.key)
	  link.Reverse()
	  fmt.Println("\n==============================\n")*/
}
