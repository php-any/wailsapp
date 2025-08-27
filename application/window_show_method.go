package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowShowMethod struct {
	source applicationsrc.Window
}

func (h *WindowShowMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Show()
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowShowMethod) GetName() string            { return "show" }
func (h *WindowShowMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowShowMethod) GetIsStatic() bool          { return true }
func (h *WindowShowMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowShowMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowShowMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
