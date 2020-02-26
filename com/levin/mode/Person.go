package mode

import "fmt"

type Person struct {
	name    string
	age     int
	country Country
}
type Country int

const (
	China Country = 1 << iota
	India
	America
)

func (p *Person) SayHello() {
	fmt.Printf("Hello,I'am %s\n", p.name)
}
