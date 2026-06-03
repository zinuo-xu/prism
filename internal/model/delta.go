package model

type DeltaType int

const (
	Added       DeltaType = iota
	Removed
	Changed
	Reordered
	TypeChanged
)

type Delta struct {
	Path     string
	Type     DeltaType
	OldValue interface{}
	NewValue interface{}
	Message  string
}
