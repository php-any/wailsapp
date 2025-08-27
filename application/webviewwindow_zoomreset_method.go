package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowZoomResetMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowZoomResetMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.ZoomReset()
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowZoomResetMethod) GetName() string            { return "zoomReset" }
func (h *WebviewWindowZoomResetMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowZoomResetMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowZoomResetMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowZoomResetMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowZoomResetMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
