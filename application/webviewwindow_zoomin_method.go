package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowZoomInMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowZoomInMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ZoomIn()
	return nil, nil
}

func (h *WebviewWindowZoomInMethod) GetName() string            { return "zoomIn" }
func (h *WebviewWindowZoomInMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowZoomInMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowZoomInMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowZoomInMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowZoomInMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
