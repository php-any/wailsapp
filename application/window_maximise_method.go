package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowMaximiseMethod struct {
	source applicationsrc.Window
}

func (h *WindowMaximiseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Maximise()
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowMaximiseMethod) GetName() string            { return "maximise" }
func (h *WindowMaximiseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowMaximiseMethod) GetIsStatic() bool          { return true }
func (h *WindowMaximiseMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowMaximiseMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowMaximiseMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
