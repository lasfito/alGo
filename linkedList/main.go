package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head *node
	tail *node
}

func (l *linkedList) addNode(value int) {

	newNode := new(node)
	newNode.value = value

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail = newNode
		current := l.head
		// no necesitamos esto al usar tail, lol
		for ; current.next != nil; current = current.next {
		}
		current.next = newNode

	}
}

func (l *linkedList) String() string {
	// ¿por qué llaves y no paréntesis?
	sb := strings.Builder{}

	for current := l.head; current != nil; current = current.next {
		// putos nombres de format xd, demennnn mi console.log :,v o String()
		// investigar estas madres e implementar método string en node
		sb.WriteString(fmt.Sprintf("%d-", current.value))
	}

	return sb.String()
}

func (l *linkedList) delete(value int) {

	// ¿por qué *node y no solo node?
	var previous *node

	if value == l.head.value {
		l.head = l.head.next
	} else {

		for current := l.head; current != nil; current = current.next {
			if (current.value == value) && (previous != nil) {
				previous.next = current.next
			}
			if (current.value == value) && (previous == nil) {
				l.head = current.next
			}

			previous = current

		}

	}
}

func main() {

	// este modo de inicializar
	list := new(linkedList)
	// vs list:= linkedList{} ???
	list.addNode(5)
	list.addNode(10)
	list.addNode(15)
	list.addNode(20)

	fmt.Println(list)

	list.delete(10)
	fmt.Println(list)

}
