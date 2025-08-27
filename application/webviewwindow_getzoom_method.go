package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowGetZoomMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowGetZoomMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetZoom()
	return data.NewAnyValue(ret0), nil
}

func (h *WebviewWindowGetZoomMethod) GetName() string            { return "getZoom" }
func (h *WebviewWindowGetZoomMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowGetZoomMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowGetZoomMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WebviewWindowGetZoomMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WebviewWindowGetZoomMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
