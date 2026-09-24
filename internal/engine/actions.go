package engine

type Action interface {
	IsAction() bool
}

type EscapeAction struct {
}

func (e EscapeAction) IsAction() bool {
	return true
}
