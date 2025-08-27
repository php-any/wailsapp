package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowGetZoomMethod struct {
	source applicationsrc.Window
}

func (h *WindowGetZoomMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.GetZoom()
	return data.NewAnyValue(ret0), nil
}

func (h *WindowGetZoomMethod) GetName() string            { return "getZoom" }
func (h *WindowGetZoomMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowGetZoomMethod) GetIsStatic() bool          { return true }
func (h *WindowGetZoomMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowGetZoomMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowGetZoomMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
