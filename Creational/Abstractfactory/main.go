package main

import ui "github.com/suleiman-oss/Design_Patterns_in_Go/Creational/Abstractfactory/UI"

func main() {
	//Android UI
	android := ui.AndroidUI{}
	android.SetColor()
	android.SetTheme()
	androidFactory := android.CreateFactory()
	androidMenu := androidFactory.GetMenu()
	androidMenu.ShowMenu()
	androidButton := androidFactory.GetButton()
	androidButton.ClickButton()
	//IOS UI
	ios := ui.IOSUI{}
	ios.SetColor()
	ios.SetTheme()
	iosFactory := ios.CreateFactory()
	iosMenu := iosFactory.GetMenu()
	iosMenu.ShowMenu()
	iosButton := iosFactory.GetButton()
	iosButton.ClickButton()
}
