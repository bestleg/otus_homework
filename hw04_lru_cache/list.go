package hw04lrucache

import "log"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

//// Next returns the next list element or nil.
//func (e *ListItem) Next() *ListItem {
//	if p := e.next; e.list != nil && p != &e.list.root {
//		return p
//	}
//	return nil
//}
//
//// Prev returns the previous list element or nil.
//func (e *ListItem) Prev() *ListItem {
//	if p := e.prev; e.list != nil && p != &e.list.root {
//		return p
//	}
//	return nil
//}

type list struct {
	root  ListItem
	first ListItem
	last  ListItem
	len   int
}

func (l *list) Init() *list {
	l.root.Next = &l.root
	l.root.Prev = &l.root
	l.len = 0
	return l
}

func NewList() *list { return new(list).Init() }

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.root.Next
}

func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.root.Prev
}

func (l *list) PushFront(v interface{}) *ListItem {
	newFirstItem := &ListItem{
		Value: v,
	}
	newFirstItem.Prev = &l.root
	newFirstItem.Next = l.root.Next
	newFirstItem.Prev.Next = newFirstItem
	newFirstItem.Next.Prev = newFirstItem
	l.len++
	log.Println("front", newFirstItem)
	return newFirstItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newLastItem := &ListItem{
		Value: v,
	}
	newLastItem.Prev = l.root.Prev
	newLastItem.Next = &l.root
	newLastItem.Prev.Next = newLastItem
	newLastItem.Next.Prev = newLastItem
	l.len++
	log.Println("back", newLastItem)
	return newLastItem
}

func (l *list) Remove(i *ListItem) {
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	i.Next = nil // avoid memory leaks
	i.Prev = nil // avoid memory leaks
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	i.Next = l.root.Next
	l.root.Next = i
}
