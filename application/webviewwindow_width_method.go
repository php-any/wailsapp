package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowWidthMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowWidthMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Width()
	return data.NewIntValue(ret0), nil
}

func (h *WebviewWindowWidthMethod) GetName() string            { return "width" }
func (h *WebviewWindowWidthMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowWidthMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowWidthMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowWidthMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowWidthMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
