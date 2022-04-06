package main

import (
	"fmt"
)

type List struct {
	Data string
	Next *List
}

//need to be able to add, remove, and traverse a list



// adds an element after the current element
func (L *List) addNext(next string){
	if L.Data == "" {
		L.Data = next
		return
	}
	temp := List{
		Data: next,
		Next: L.Next,
	}
	L.Next = &temp
}

// adds an element to the end of a list
func (L *List) add(next string){
	if L.Data == "" {
		L.Data = next
		return
	}
	elt := L
	for {
		if elt.Next == nil {
			temp := List{
				Data: next,
			}
			elt.Next = &temp
			break
		} else {
			elt = L.Next
		}
	}
}

// removes the next element
func (L *List) removeNext(){
	if L.Next == nil {
		//there is no next item to remove
		return
	}
	L.Next = L.Next.Next
}

// removes the last element
func (L *List) removeLast(){
	if L.Next == nil{
		L = nil
		return //is already the last item in the list
	}
	elt := L
	for {
		if elt.Next.Next == nil{
			elt.Next = nil
			break
		}
		elt = elt.Next

	}
}

//prints out all elements of the list
func (L *List) printList() {
	elt := L
	for elt != nil {
		fmt.Printf("%v\n",elt.Data)
		elt = elt.Next
	}
}

// removes specified element 
func (L *List) remove(item string){
	if L.Data == item{
		if L.Next != nil{
			*L = *L.Next
		}else{
			L.Data = ""
		}

		return
	}
	elt := L 
	for {
		
		if elt.Next == nil{
			fmt.Println("this item is not in the list")
			return
		}
		if elt.Next.Data == item{
			elt.Next = elt.Next.Next
			return
		}
	}

}

// traverses the list
// traverses to a specific element in the list