package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowIsFullscreenMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowIsFullscreenMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsFullscreen()
	return data.NewBoolValue(ret0), nil
}

func (h *WebviewWindowIsFullscreenMethod) GetName() string            { return "isFullscreen" }
func (h *WebviewWindowIsFullscreenMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowIsFullscreenMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowIsFullscreenMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowIsFullscreenMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowIsFullscreenMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
