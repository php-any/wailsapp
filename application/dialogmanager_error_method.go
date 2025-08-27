package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type DialogManagerErrorMethod struct {
	source *applicationsrc.DialogManager
}

func (h *DialogManagerErrorMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Error()
	return data.NewClassValue(NewMessageDialogClassFrom(ret0), ctx), nil
}

func (h *DialogManagerErrorMethod) GetName() string            { return "error" }
func (h *DialogManagerErrorMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *DialogManagerErrorMethod) GetIsStatic() bool          { return true }
func (h *DialogManagerErrorMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *DialogManagerErrorMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *DialogManagerErrorMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
