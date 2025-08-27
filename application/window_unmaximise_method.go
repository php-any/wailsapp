package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowUnMaximiseMethod struct {
	source applicationsrc.Window
}

func (h *WindowUnMaximiseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.UnMaximise()
	return nil, nil
}

func (h *WindowUnMaximiseMethod) GetName() string            { return "unMaximise" }
func (h *WindowUnMaximiseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowUnMaximiseMethod) GetIsStatic() bool          { return true }
func (h *WindowUnMaximiseMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowUnMaximiseMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowUnMaximiseMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
