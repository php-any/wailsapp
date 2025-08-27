package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowPositionMethod struct {
	source applicationsrc.Window
}

func (h *WindowPositionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0, ret1 := h.source.Position()
	return data.NewAnyValue([]any{ret0, ret1}), nil
}

func (h *WindowPositionMethod) GetName() string            { return "position" }
func (h *WindowPositionMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowPositionMethod) GetIsStatic() bool          { return true }
func (h *WindowPositionMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowPositionMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowPositionMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
