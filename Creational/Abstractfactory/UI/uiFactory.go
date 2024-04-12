package ui

type UIFactory interface {
	GetMenu() Menu
	GetButton() Button
}

type AndroidFactory struct{}

func (af *AndroidFactory) GetMenu() Menu {
	return &AndroidMenu{}
}

func (af *AndroidFactory) GetButton() Button {
	return &AndoridButton{}
}

type IOSFactory struct{}

func (i *IOSFactory) GetMenu() Menu {
	return &IOSMenu{}
}

func (i *IOSFactory) GetButton() Button {
	return &IOSButton{}
}
