package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowRelativePositionMethod struct {
	source applicationsrc.Window
}

func (h *WindowRelativePositionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0, ret1 := h.source.RelativePosition()
	return data.NewAnyValue([]any{ret0, ret1}), nil
}

func (h *WindowRelativePositionMethod) GetName() string            { return "relativePosition" }
func (h *WindowRelativePositionMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowRelativePositionMethod) GetIsStatic() bool          { return true }
func (h *WindowRelativePositionMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowRelativePositionMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowRelativePositionMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
