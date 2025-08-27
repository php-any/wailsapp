package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowOpenDevToolsMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowOpenDevToolsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.OpenDevTools()
	return nil, nil
}

func (h *WebviewWindowOpenDevToolsMethod) GetName() string            { return "openDevTools" }
func (h *WebviewWindowOpenDevToolsMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowOpenDevToolsMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowOpenDevToolsMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowOpenDevToolsMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowOpenDevToolsMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
