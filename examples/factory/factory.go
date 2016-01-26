package factory

import "fmt"

type Button interface {
	Paint()
	OnClick()
}

type Label interface {
	Paint()
}

// WinButton is a Button implementation for Windows.
type WinButton struct{}

func (WinButton) Paint()   { fmt.Println("win button paint") }
func (WinButton) OnClick() { fmt.Println("win button click") }

// WinLabel is a Label implementation for Windows.
type WinLabel struct{}

func (WinLabel) Paint() { fmt.Println("win label paint") }

// WinButton is a Button implementation for Mac.
type MacButton struct{}

func (MacButton) Paint()   { fmt.Println("mac button paint") }
func (MacButton) OnClick() { fmt.Println("mac button click") }

// WinLabel is a Label implementation for Mac.
type MacLabel struct{}

func (MacLabel) Paint() { fmt.Println("mac label paint") }

// UI factory can create buttons and labels.
type UIFactory interface {
	CreateButton() Button
	CreateLabel() Label
}

// WinFactory is a UI factory that can create Windows UI elements.
type WinFactory struct{}

func (WinFactory) CreateButton() Button {
	return WinButton{}
}

func (WinFactory) CreateLabel() Label {
	return WinLabel{}
}

// MacFactory is a UI factory that can create Mac UI elements.
type MacFactory struct{}

func (MacFactory) CreateButton() Button {
	return MacButton{}
}

func (MacFactory) CreateLabel() Label {
	return MacLabel{}
}

// CreateFactory returns a UIFactory of the given os.
func CreateFactory(os string) UIFactory {
	if os == "win" {
		return WinFactory{}
	} else {
		return MacFactory{}
	}
}

func Run(f UIFactory) {
	button := f.CreateButton()
	button.Paint()
	button.OnClick()
	label := f.CreateLabel()
	label.Paint()
}
