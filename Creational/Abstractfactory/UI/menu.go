package ui

import "fmt"

type Menu interface {
	ShowMenu()
}

type AndroidMenu struct{}

func (am *AndroidMenu) ShowMenu() {
	fmt.Println("Showing Android Menu")
}

type IOSMenu struct{}

func (im *IOSMenu) ShowMenu() {
	fmt.Println("Showing IOS Menu")
}
