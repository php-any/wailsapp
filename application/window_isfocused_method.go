package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowIsFocusedMethod struct {
	source applicationsrc.Window
}

func (h *WindowIsFocusedMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsFocused()
	return data.NewBoolValue(ret0), nil
}

func (h *WindowIsFocusedMethod) GetName() string            { return "isFocused" }
func (h *WindowIsFocusedMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowIsFocusedMethod) GetIsStatic() bool          { return true }
func (h *WindowIsFocusedMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowIsFocusedMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowIsFocusedMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
