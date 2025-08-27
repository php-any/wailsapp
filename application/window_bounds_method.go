package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowBoundsMethod struct {
	source applicationsrc.Window
}

func (h *WindowBoundsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Bounds()
	return data.NewClassValue(NewRectClassFrom(&ret0), ctx), nil
}

func (h *WindowBoundsMethod) GetName() string            { return "bounds" }
func (h *WindowBoundsMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowBoundsMethod) GetIsStatic() bool          { return true }
func (h *WindowBoundsMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowBoundsMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowBoundsMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
