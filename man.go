package fine

type Man struct {
	Name    string
	Age     int
	Address string
}

func (man *Man) SayName() string {
	return man.Name
}

func (man *Man) SayAge() int {
	return man.Age
}

func (man *Man) SayAddress() string {
	return man.Address
}