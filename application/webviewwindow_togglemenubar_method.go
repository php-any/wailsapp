package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowToggleMenuBarMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowToggleMenuBarMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ToggleMenuBar()
	return nil, nil
}

func (h *WebviewWindowToggleMenuBarMethod) GetName() string            { return "toggleMenuBar" }
func (h *WebviewWindowToggleMenuBarMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowToggleMenuBarMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowToggleMenuBarMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowToggleMenuBarMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowToggleMenuBarMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
