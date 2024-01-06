package chapter3

type LinkedList struct {
	item any
	next *LinkedList
}

var NULL = LinkedList{
	item: nil,
	next: nil,
}

func search(item any, list *LinkedList) any {
	if list == nil {
		return nil
	}

	if list.item == item {
		return list
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

func remove(doomedNode *LinkedList, list *LinkedList) {
	pred := predecessor(doomedNode, list)
	if pred == nil {
		return
	} else {
		pred.next = doomedNode.next
	}
}

func predecessor(doomedNode *LinkedList, origList *LinkedList) *LinkedList {
	if origList == nil || origList.next == nil {
		return &LinkedList{}
	}

	if origList.next == doomedNode {
		return origList
	} else {
		return predecessor(doomedNode, origList.next)
	}
}
