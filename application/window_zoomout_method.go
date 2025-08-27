package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowZoomOutMethod struct {
	source applicationsrc.Window
}

func (h *WindowZoomOutMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.ZoomOut()
	return nil, nil
}

func (h *WindowZoomOutMethod) GetName() string            { return "zoomOut" }
func (h *WindowZoomOutMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowZoomOutMethod) GetIsStatic() bool          { return true }
func (h *WindowZoomOutMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowZoomOutMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowZoomOutMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
