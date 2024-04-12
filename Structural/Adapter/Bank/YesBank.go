package bank

type YesbankAPI interface {
	GetBalance() string
}

type YesBank struct{}

func (y *YesBank) GetBalance() string {
	return "Your balance in $1000"
}
