package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowUnFullscreenMethod struct {
	source applicationsrc.Window
}

func (h *WindowUnFullscreenMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.UnFullscreen()
	return nil, nil
}

func (h *WindowUnFullscreenMethod) GetName() string            { return "unFullscreen" }
func (h *WindowUnFullscreenMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowUnFullscreenMethod) GetIsStatic() bool          { return true }
func (h *WindowUnFullscreenMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowUnFullscreenMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowUnFullscreenMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
