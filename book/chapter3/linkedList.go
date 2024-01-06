package chapter3

type LinkedList struct {
	item any
	next *LinkedList
}

func searchList(item any, list *LinkedList) any {
	if list == nil {
		return nil
	}

	if list.item == item {
		return item
	} else {
		return searchList(item, list.next)
	}
}
