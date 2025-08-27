package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowResizableMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowResizableMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Resizable()
	return data.NewBoolValue(ret0), nil
}

func (h *WebviewWindowResizableMethod) GetName() string            { return "resizable" }
func (h *WebviewWindowResizableMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowResizableMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowResizableMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowResizableMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowResizableMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
