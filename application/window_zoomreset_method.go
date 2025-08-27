package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowZoomResetMethod struct {
	source applicationsrc.Window
}

func (h *WindowZoomResetMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.ZoomReset()
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowZoomResetMethod) GetName() string            { return "zoomReset" }
func (h *WindowZoomResetMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowZoomResetMethod) GetIsStatic() bool          { return true }
func (h *WindowZoomResetMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowZoomResetMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowZoomResetMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
