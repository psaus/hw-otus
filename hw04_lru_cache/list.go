package hw04lrucache

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

type list struct {
	frontItemPtr *ListItem
	backItempPtr *ListItem
	len          int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.frontItemPtr
}

func (l *list) Back() *ListItem {
	return l.backItempPtr
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	l.len++

	if l.frontItemPtr == nil && l.backItempPtr == nil {
		l.frontItemPtr = item
		l.backItempPtr = item
		return item
	}

	if l.frontItemPtr != nil {
		item.Next = l.frontItemPtr
		l.frontItemPtr.Prev = item
	}

	l.frontItemPtr = item
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{Value: v}
	l.len++

	if l.frontItemPtr == nil && l.backItempPtr == nil {
		l.frontItemPtr = item
		l.backItempPtr = item
		return item
	}

	if l.backItempPtr != nil {
		item.Prev = l.backItempPtr
		l.backItempPtr.Next = item
	}

	l.backItempPtr = item
	return item
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	// get linked items
	prevItem, nextItem := i.Prev, i.Next

	switch {
	case prevItem != nil && nextItem != nil:
		// the item to be removed has middle position
		prevItem.Next, nextItem.Prev = nextItem, prevItem
	case prevItem == nil:
		// the item to be removed has front position
		nextItem.Prev, l.frontItemPtr = nil, nextItem
	case nextItem == nil:
		// the item to be removed has back position
		prevItem.Next, l.backItempPtr = nil, prevItem
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	// skip invalid item
	if l.frontItemPtr == nil || i.Prev == nil || l.frontItemPtr == i {
		return
	}

	// link elements between the current
	if i.Next != nil {
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	} else {
		i.Prev.Next = nil
	}

	// move first element to next
	currentFrontItem := l.frontItemPtr
	currentFrontItem.Prev = i

	// set current item to front
	i.Next = currentFrontItem
	i.Prev = nil
	l.frontItemPtr = i
}

func NewList() List {
	return new(list)
}
