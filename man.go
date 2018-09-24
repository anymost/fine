package fine

import "strconv"

type Man struct {
	Name    string
	Age     int
	Address string
	Prev, Next *Man
}

func (man *Man) String() string {
	return man.Name + " " + strconv.Itoa(man.Age) + " " + man.Address
}
