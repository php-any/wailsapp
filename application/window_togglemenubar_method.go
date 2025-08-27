package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowToggleMenuBarMethod struct {
	source applicationsrc.Window
}

func (h *WindowToggleMenuBarMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ToggleMenuBar()
	return nil, nil
}

func (h *WindowToggleMenuBarMethod) GetName() string            { return "toggleMenuBar" }
func (h *WindowToggleMenuBarMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowToggleMenuBarMethod) GetIsStatic() bool          { return true }
func (h *WindowToggleMenuBarMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowToggleMenuBarMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowToggleMenuBarMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
