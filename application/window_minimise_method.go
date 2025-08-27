package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowMinimiseMethod struct {
	source applicationsrc.Window
}

func (h *WindowMinimiseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Minimise()
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowMinimiseMethod) GetName() string            { return "minimise" }
func (h *WindowMinimiseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowMinimiseMethod) GetIsStatic() bool          { return true }
func (h *WindowMinimiseMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowMinimiseMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowMinimiseMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
