package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowToggleFullscreenMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowToggleFullscreenMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ToggleFullscreen()
	return nil, nil
}

func (h *WebviewWindowToggleFullscreenMethod) GetName() string            { return "toggleFullscreen" }
func (h *WebviewWindowToggleFullscreenMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowToggleFullscreenMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowToggleFullscreenMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowToggleFullscreenMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowToggleFullscreenMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
