package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowZoomMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowZoomMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Zoom()
	return nil, nil
}

func (h *WebviewWindowZoomMethod) GetName() string            { return "zoom" }
func (h *WebviewWindowZoomMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowZoomMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowZoomMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowZoomMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowZoomMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
