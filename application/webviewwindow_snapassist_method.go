package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSnapAssistMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSnapAssistMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.SnapAssist()
	return nil, nil
}

func (h *WebviewWindowSnapAssistMethod) GetName() string            { return "snapAssist" }
func (h *WebviewWindowSnapAssistMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowSnapAssistMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowSnapAssistMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowSnapAssistMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowSnapAssistMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
