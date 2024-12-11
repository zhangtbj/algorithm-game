package main

import "fmt"

func main() {
	list := Constructor()
	list.AddAtHead(1)
	//list.AddAtTail(1)
	//list.AddAtTail(2)
	list.AddAtTail(3)
	//list.AddAtHead(1)
	//fmt.Println(list.Get(0))
	//fmt.Println(list.Get(1))
	//fmt.Println(list.Get(2))
	list.AddAtIndex(1, 2)
	fmt.Println(list.Get(1))
	list.DeleteAtIndex(1)
	fmt.Println(list.Get(1))
	fmt.Println(list.Get(3))
	list.DeleteAtIndex(3)
	list.DeleteAtIndex(0)
	fmt.Println(list.Get(0))
	list.DeleteAtIndex(0)
	fmt.Println(list.Get(0))
	list.Get(0)
	list.printLinkedList()
}

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{Val: -1}
}

func (this *MyLinkedList) Get(index int) int {
	current := this
	for i := 0; current != nil; i++ {
		if i == index {
			return current.Val
		} else {
			current = current.Next
		}
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.Val == -1 {
		this.Val = val
	} else {
		tmp := &MyLinkedList{Val: this.Val, Next: this.Next}
		this.Val = val
		this.Next = tmp
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	//current := this
	if this.Val == -1 {
		this.Val = val
	} else {
		newNode := &MyLinkedList{Val: val}

		for this.Next != nil {
			this = this.Next
		}

		this.Next = newNode
	}

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	current := this
	for i := 0; current != nil; i++ {
		if i == index-1 {
			break
		}
		current = current.Next
	}

	if current != nil {
		newNode := &MyLinkedList{Val: val, Next: current.Next}
		current.Next = newNode
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if this.Next != nil {
			this.Val = this.Next.Val
			this.Next = this.Next.Next
		} else {
			this.Val = -1
		}
		return
	}

	current := this
	for i := 0; current != nil; i++ {
		if i == index-1 {
			break
		}
		current = current.Next
	}

	if current != nil && current.Next != nil {
		current.Next = current.Next.Next
	}
}

// 打印链表
func (this *MyLinkedList) printLinkedList() {
	current := this
	for current != nil { // 遍历链表
		fmt.Printf("%d->", current.Val) // 打印节点值
		current = current.Next          // 切换到下一个节点
	}
	fmt.Println("NULL")
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
