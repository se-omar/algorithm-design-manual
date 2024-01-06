package chapter3

type LinkedList struct {
	item any
	next *LinkedList
}

func search(item any, list *LinkedList) any {
	if list == nil {
		return nil
	}

	if list.item == item {
		return item
	} else {
		return search(item, list.next)
	}
}

func insert(item any, list *LinkedList) {
	newL := LinkedList{
		item: item,
		next: list,
	}

	*list = newL
}
