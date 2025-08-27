package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowNativeWindowMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowNativeWindowMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.NativeWindow()
	return data.NewAnyValue(ret0), nil
}

func (h *WebviewWindowNativeWindowMethod) GetName() string            { return "nativeWindow" }
func (h *WebviewWindowNativeWindowMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowNativeWindowMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowNativeWindowMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowNativeWindowMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowNativeWindowMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
