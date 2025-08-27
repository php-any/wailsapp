package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowToggleMaximiseMethod struct {
	source applicationsrc.Window
}

func (h *WindowToggleMaximiseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ToggleMaximise()
	return nil, nil
}

func (h *WindowToggleMaximiseMethod) GetName() string            { return "toggleMaximise" }
func (h *WindowToggleMaximiseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowToggleMaximiseMethod) GetIsStatic() bool          { return true }
func (h *WindowToggleMaximiseMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowToggleMaximiseMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowToggleMaximiseMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
