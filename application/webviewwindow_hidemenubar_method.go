package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowHideMenuBarMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowHideMenuBarMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.HideMenuBar()
	return nil, nil
}

func (h *WebviewWindowHideMenuBarMethod) GetName() string            { return "hideMenuBar" }
func (h *WebviewWindowHideMenuBarMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowHideMenuBarMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowHideMenuBarMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowHideMenuBarMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowHideMenuBarMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
