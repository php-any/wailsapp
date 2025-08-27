package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowEventContextMethod struct {
	source *applicationsrc.WindowEvent
}

func (h *WindowEventContextMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Context()
	return data.NewClassValue(NewWindowEventContextClassFrom(ret0), ctx), nil
}

func (h *WindowEventContextMethod) GetName() string            { return "context" }
func (h *WindowEventContextMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowEventContextMethod) GetIsStatic() bool          { return true }
func (h *WindowEventContextMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowEventContextMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowEventContextMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
