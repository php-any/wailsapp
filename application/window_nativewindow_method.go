package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowNativeWindowMethod struct {
	source applicationsrc.Window
}

func (h *WindowNativeWindowMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.NativeWindow()
	return data.NewAnyValue(ret0), nil
}

func (h *WindowNativeWindowMethod) GetName() string            { return "nativeWindow" }
func (h *WindowNativeWindowMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowNativeWindowMethod) GetIsStatic() bool          { return true }
func (h *WindowNativeWindowMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowNativeWindowMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowNativeWindowMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
