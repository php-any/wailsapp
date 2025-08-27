package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowSizeMethod struct {
	source applicationsrc.Window
}

func (h *WindowSizeMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0, ret1 := h.source.Size()
	return data.NewAnyValue([]any{ret0, ret1}), nil
}

func (h *WindowSizeMethod) GetName() string            { return "size" }
func (h *WindowSizeMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowSizeMethod) GetIsStatic() bool          { return true }
func (h *WindowSizeMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowSizeMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowSizeMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
