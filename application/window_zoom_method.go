package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowZoomMethod struct {
	source applicationsrc.Window
}

func (h *WindowZoomMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Zoom()
	return nil, nil
}

func (h *WindowZoomMethod) GetName() string            { return "zoom" }
func (h *WindowZoomMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowZoomMethod) GetIsStatic() bool          { return true }
func (h *WindowZoomMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowZoomMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowZoomMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
