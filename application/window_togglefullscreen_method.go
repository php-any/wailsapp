package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowToggleFullscreenMethod struct {
	source applicationsrc.Window
}

func (h *WindowToggleFullscreenMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ToggleFullscreen()
	return nil, nil
}

func (h *WindowToggleFullscreenMethod) GetName() string            { return "toggleFullscreen" }
func (h *WindowToggleFullscreenMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowToggleFullscreenMethod) GetIsStatic() bool          { return true }
func (h *WindowToggleFullscreenMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowToggleFullscreenMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowToggleFullscreenMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
