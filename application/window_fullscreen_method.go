package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowFullscreenMethod struct {
	source applicationsrc.Window
}

func (h *WindowFullscreenMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Fullscreen()
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowFullscreenMethod) GetName() string            { return "fullscreen" }
func (h *WindowFullscreenMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowFullscreenMethod) GetIsStatic() bool          { return true }
func (h *WindowFullscreenMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowFullscreenMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowFullscreenMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
