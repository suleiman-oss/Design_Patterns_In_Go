package main

import bank "github.com/suleiman-oss/Design_Patterns_in_Go/Structural/Adapter/Bank"

func main() {
	y := &bank.YesBank{}
	adp := bank.NewBankAdapter(y)
	p := bank.NewPhonePe(adp)
	p.ShowBalance()
}
