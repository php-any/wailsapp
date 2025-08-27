package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowFullscreenMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowFullscreenMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Fullscreen()
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowFullscreenMethod) GetName() string            { return "fullscreen" }
func (h *WebviewWindowFullscreenMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowFullscreenMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowFullscreenMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowFullscreenMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowFullscreenMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
