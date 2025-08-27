package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowShowMenuBarMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowShowMenuBarMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ShowMenuBar()
	return nil, nil
}

func (h *WebviewWindowShowMenuBarMethod) GetName() string            { return "showMenuBar" }
func (h *WebviewWindowShowMenuBarMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowShowMenuBarMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowShowMenuBarMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowShowMenuBarMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowShowMenuBarMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
