type LinkedList struct {
	head, tail *node
	leng       int
}

type node struct {
	value int
	nxt   *node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	if index < 0 || index >= ll.leng {
		return -1
	}
	ctr := 0
	ptr := ll.head
	for index != ctr {
		ptr = ptr.nxt
		ctr++
	}

	return ptr.value
}

func (ll *LinkedList) InsertHead(val int) {
	nde := new(node)
	nde.value = val
	nde.nxt = ll.head
	ll.head = nde
	if ll.tail == nil {
		ll.tail = nde
	}
	ll.leng++
}

func (ll *LinkedList) InsertTail(val int) {
	nde := new(node)
	nde.value = val
	if ll.tail == nil {
		ll.head = nde
	} else {
		ll.tail.nxt = nde
	}
	ll.tail = nde
	ll.leng++
}

func (ll *LinkedList) Remove(index int) bool {
	if index < 0 || index >= ll.leng {
		return false
	}

	if index == 0 {
		ll.head = ll.head.nxt
		if ll.head == nil {
			ll.tail = nil
		}
		ll.leng--
		return true
	}

	prev := ll.head
	for i := 0; i < index-1; i++ {
		prev = prev.nxt
	}

	curr := prev.nxt
	prev.nxt = curr.nxt
	if curr == ll.tail {
		ll.tail = prev
	}
	ll.leng--
	return true
}

func (ll *LinkedList) GetValues() []int {
	ret := make([]int, ll.leng)

	cptr := ll.head

	for i := range ret {
		ret[i] = cptr.value
		cptr = cptr.nxt
	}
	return ret
}
