package fine

import "fmt"

type ManList struct {
	Start *Man
	End *Man
	Len int
}

func MakeManList () *ManList {
	Start := &Man{
		Name: "",
		Age: 0,
		Address: "",
		Prev: nil,
		Next: nil,
	}
	End := &Man{
		Name: "",
		Age: 0,
		Address: "",
		Prev: nil,
		Next: nil,
	}
	Start.Next = End
	End.Prev = Start
	return &ManList{
		Start,
		End,
		0,
	}
}

func (list *ManList) Push(item *Man)  {
	list.End.Prev.Next = item
	list.End.Prev = item
	list.Len += 1
}

func (list *ManList) UnShift(item *Man)  {
	list.Start.Next.Prev = item
	list.Start.Next = item
	list.Len += 1
}

func (list *ManList) Traverse()  {
	start := list.Start
	for {
		if start.Next == nil {
			return
		}
		fmt.Println(start.Next)
		start = start.Next
	}
}