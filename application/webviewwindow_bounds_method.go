package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowBoundsMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowBoundsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Bounds()
	return data.NewClassValue(NewRectClassFrom(&ret0), ctx), nil
}

func (h *WebviewWindowBoundsMethod) GetName() string            { return "bounds" }
func (h *WebviewWindowBoundsMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowBoundsMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowBoundsMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowBoundsMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowBoundsMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
