package fine

type Man struct {
	Name    string
	Age     int
	Address string
}

func (man *Man) sayName() string {
	return man.Name
}

func (man *Man) sayAge() int {
	return man.Age
}

func (man *Man) sayAddress() string {
	return man.Address
}