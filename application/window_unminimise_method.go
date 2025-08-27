package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowUnMinimiseMethod struct {
	source applicationsrc.Window
}

func (h *WindowUnMinimiseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.UnMinimise()
	return nil, nil
}

func (h *WindowUnMinimiseMethod) GetName() string            { return "unMinimise" }
func (h *WindowUnMinimiseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowUnMinimiseMethod) GetIsStatic() bool          { return true }
func (h *WindowUnMinimiseMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowUnMinimiseMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowUnMinimiseMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
