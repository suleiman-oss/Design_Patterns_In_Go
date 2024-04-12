package ui

import "fmt"

type UI interface {
	SetTheme()
	SetColor()
	CreateFactory() UIFactory
}

type AndroidUI struct{}

func (a *AndroidUI) SetTheme() {
	fmt.Println("Setting theme on Android")
}

func (a *AndroidUI) SetColor() {
	fmt.Println("Setting color on Andorid")
}

func (a *AndroidUI) CreateFactory() UIFactory {
	return &AndroidFactory{}
}

type IOSUI struct{}

func (i *IOSUI) SetTheme() {
	fmt.Println("Setting theme on IOS")
}

func (i *IOSUI) SetColor() {
	fmt.Println("Setting color on IOS")
}

func (i *IOSUI) CreateFactory() UIFactory {
	return &IOSFactory{}
}
