package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowHideMenuBarMethod struct {
	source applicationsrc.Window
}

func (h *WindowHideMenuBarMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.HideMenuBar()
	return nil, nil
}

func (h *WindowHideMenuBarMethod) GetName() string            { return "hideMenuBar" }
func (h *WindowHideMenuBarMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowHideMenuBarMethod) GetIsStatic() bool          { return true }
func (h *WindowHideMenuBarMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowHideMenuBarMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowHideMenuBarMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
