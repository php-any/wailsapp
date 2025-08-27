package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowUnFullscreenMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowUnFullscreenMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.UnFullscreen()
	return nil, nil
}

func (h *WebviewWindowUnFullscreenMethod) GetName() string            { return "unFullscreen" }
func (h *WebviewWindowUnFullscreenMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowUnFullscreenMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowUnFullscreenMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowUnFullscreenMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowUnFullscreenMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
