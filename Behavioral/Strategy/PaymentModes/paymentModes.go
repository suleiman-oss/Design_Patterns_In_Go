package paymentmodes

import "fmt"

type Payment interface {
	Pay(amount float32)
}

type UPIPayment struct{}

func (u *UPIPayment) Pay(amount float32) {
	fmt.Printf("Paying amount:%f through UPI\n", amount)
}

type BankPayment struct{}

func (b *BankPayment) Pay(amount float32) {
	fmt.Printf("Paying amount:%f through Bank Transfer\n", amount)
}

type PaymentStrategy struct {
	PaymentType Payment
}

func NewPaymentStrategy(paymentType Payment) *PaymentStrategy {
	return &PaymentStrategy{PaymentType: paymentType}
}
