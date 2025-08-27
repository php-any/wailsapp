package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSizeMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSizeMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0, ret1 := h.source.Size()
	return data.NewAnyValue([]any{ret0, ret1}), nil
}

func (h *WebviewWindowSizeMethod) GetName() string            { return "size" }
func (h *WebviewWindowSizeMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowSizeMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowSizeMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowSizeMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowSizeMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
