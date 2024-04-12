package bank

import "fmt"

type PhonePeAPI interface {
	ShowBalance()
}

type BankAdapater struct {
	y *YesBank
}

func (ba *BankAdapater) ShowBalance() {
	fmt.Println(ba.y.GetBalance())
}

func NewBankAdapter(y *YesBank) *BankAdapater {
	return &BankAdapater{y: y}
}

type PhonePe struct {
	adp *BankAdapater
}

func (p *PhonePe) ShowBalance() {
	p.adp.ShowBalance()
}

func NewPhonePe(ba *BankAdapater) *PhonePe {
	return &PhonePe{adp: ba}
}
