package ui

import "fmt"

type Button interface {
	ClickButton()
}

type AndoridButton struct{}

func (ab *AndoridButton) ClickButton() {
	fmt.Println("Clicked on Android Button")
}

type IOSButton struct{}

func (ib *IOSButton) ClickButton() {
	fmt.Println("Clicked on IOS Button")
}
