package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowToggleFramelessMethod struct {
	source applicationsrc.Window
}

func (h *WindowToggleFramelessMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ToggleFrameless()
	return nil, nil
}

func (h *WindowToggleFramelessMethod) GetName() string            { return "toggleFrameless" }
func (h *WindowToggleFramelessMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowToggleFramelessMethod) GetIsStatic() bool          { return true }
func (h *WindowToggleFramelessMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowToggleFramelessMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowToggleFramelessMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
