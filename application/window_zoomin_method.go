package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowZoomInMethod struct {
	source applicationsrc.Window
}

func (h *WindowZoomInMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ZoomIn()
	return nil, nil
}

func (h *WindowZoomInMethod) GetName() string            { return "zoomIn" }
func (h *WindowZoomInMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowZoomInMethod) GetIsStatic() bool          { return true }
func (h *WindowZoomInMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowZoomInMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowZoomInMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
