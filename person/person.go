package person

import "fmt"

type Person struct {
	Age    int
	IsDead bool
}

// func (this *Person)
// func (this Person)
// *type => pointer to type
// type => concrete "type"

// func Birthday(this *Person)
// func {...person}
// func (person)

func (this *Person) Birthday() {
	fmt.Println("Before updating", this.Age)
	this.Age++
	this.IsDead = false

	fmt.Println("After updating updating", this.Age)
}

func (p *Person) DeathDay() int {
	p.IsDead = true
	return 2000
}

type HasBirthday interface {
	Birthday()
}
