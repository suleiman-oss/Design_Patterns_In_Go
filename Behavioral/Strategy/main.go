package main

import paymentmodes "github.com/suleiman-oss/Design_Patterns_in_Go/Behavioral/Strategy/PaymentModes"

func main() {
	upi := &paymentmodes.UPIPayment{}
	strategy := paymentmodes.NewPaymentStrategy(upi)
	strategy.PaymentType.Pay(12.25)

	bank := &paymentmodes.BankPayment{}
	strategy.PaymentType = bank
	strategy.PaymentType.Pay(10.05)
}
