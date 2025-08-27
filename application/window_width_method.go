package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowWidthMethod struct {
	source applicationsrc.Window
}

func (h *WindowWidthMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.Width()
	return data.NewIntValue(ret0), nil
}

func (h *WindowWidthMethod) GetName() string            { return "width" }
func (h *WindowWidthMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowWidthMethod) GetIsStatic() bool          { return true }
func (h *WindowWidthMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowWidthMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowWidthMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
