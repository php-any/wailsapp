package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowIsFocusedMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowIsFocusedMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsFocused()
	return data.NewBoolValue(ret0), nil
}

func (h *WebviewWindowIsFocusedMethod) GetName() string            { return "isFocused" }
func (h *WebviewWindowIsFocusedMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowIsFocusedMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowIsFocusedMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowIsFocusedMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowIsFocusedMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
