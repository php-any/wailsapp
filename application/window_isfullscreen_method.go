package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowIsFullscreenMethod struct {
	source applicationsrc.Window
}

func (h *WindowIsFullscreenMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsFullscreen()
	return data.NewBoolValue(ret0), nil
}

func (h *WindowIsFullscreenMethod) GetName() string            { return "isFullscreen" }
func (h *WindowIsFullscreenMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowIsFullscreenMethod) GetIsStatic() bool          { return true }
func (h *WindowIsFullscreenMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowIsFullscreenMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowIsFullscreenMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
