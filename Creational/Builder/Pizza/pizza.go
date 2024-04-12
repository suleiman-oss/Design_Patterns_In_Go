package pizza

type pizza struct {
	size    string
	crust   string
	topping string
}

func newPizza(size string, crust string, topping string) *pizza {
	return &pizza{size: size, crust: crust, topping: topping}
}
func (p *pizza) GetSize() string {
	return p.size
}

func (p *pizza) GetCrust() string {
	return p.crust
}

func (p *pizza) GetToppings() string {
	return p.topping
}

type PizzaBuilder struct {
	size    string
	crust   string
	topping string
}

func (pb *PizzaBuilder) SetSize(s string) *PizzaBuilder {
	pb.size = s
	return pb
}

func (pb *PizzaBuilder) SetCrust(c string) *PizzaBuilder {
	pb.crust = c
	return pb
}

func (pb *PizzaBuilder) AddToppings(t string) *PizzaBuilder {
	pb.topping = t
	return pb
}

func (pb *PizzaBuilder) Build() *pizza {
	//validations of the parameters
	p := newPizza(pb.size, pb.crust, pb.topping)
	return p
}

func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{}
}
