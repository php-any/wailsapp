package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowShowMenuBarMethod struct {
	source applicationsrc.Window
}

func (h *WindowShowMenuBarMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ShowMenuBar()
	return nil, nil
}

func (h *WindowShowMenuBarMethod) GetName() string            { return "showMenuBar" }
func (h *WindowShowMenuBarMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowShowMenuBarMethod) GetIsStatic() bool          { return true }
func (h *WindowShowMenuBarMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowShowMenuBarMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowShowMenuBarMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
