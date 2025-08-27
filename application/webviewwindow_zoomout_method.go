package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowZoomOutMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowZoomOutMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ZoomOut()
	return nil, nil
}

func (h *WebviewWindowZoomOutMethod) GetName() string            { return "zoomOut" }
func (h *WebviewWindowZoomOutMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowZoomOutMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowZoomOutMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowZoomOutMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowZoomOutMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
