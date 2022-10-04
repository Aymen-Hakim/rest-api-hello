package person

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func New(first, last string, age int) *Person {
	return &Person{first, last, age}
}
