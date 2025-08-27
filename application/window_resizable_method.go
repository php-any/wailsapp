package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowResizableMethod struct {
	source applicationsrc.Window
}

func (h *WindowResizableMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Resizable()
	return data.NewBoolValue(ret0), nil
}

func (h *WindowResizableMethod) GetName() string            { return "resizable" }
func (h *WindowResizableMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowResizableMethod) GetIsStatic() bool          { return true }
func (h *WindowResizableMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowResizableMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowResizableMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
