package app

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type MessageInterface interface {
	Message() string
	ActionStatus(msg string)
}

type MessageStruct struct {
	msgStatus widget.Label
}

func (m MessageStruct) Message() *widget.Label {
	return m.msgStatus
}

func (m *MessageStruct) ActionStatus(msg string) {
	str := binding.NewString()
	str.Set("")

	text := widget.NewLabelWithData(str)

	go func() {
		str.Set(msg)
	}()

	text = widget.NewLabel(text)
	m.msgStatus = text
}
