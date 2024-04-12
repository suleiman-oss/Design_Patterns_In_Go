package icecream

type Base interface {
	getCost() int
	getDescription() string
}
type cone struct {
	price int
	base  Base
}

func NewCone() *cone {
	return &cone{price: 10}
}
func NewConeWithBase(base Base) *cone {
	return &cone{price: 10, base: base}
}
func (c *cone) getCost() int {
	if c.base == nil {
		return c.price
	}
	return c.base.getCost() + c.price
}
func (c *cone) getDescription() string {
	if c.base == nil {
		return "This is a cone\n"
	}
	return c.base.getDescription() + "This is a cone\n"
}

type cream struct {
	price int
	base  Base
}

func NewCream(base Base) *cream {
	return &cream{price: 5, base: base}
}

func (c *cream) GetCost() int {
	return c.base.getCost() + c.price
}
func (c *cream) GetDescription() string {
	return c.base.getDescription() + "This is a cream\n"
}
