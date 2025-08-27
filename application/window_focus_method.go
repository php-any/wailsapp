package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowFocusMethod struct {
	source applicationsrc.Window
}

func (h *WindowFocusMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Focus()
	return nil, nil
}

func (h *WindowFocusMethod) GetName() string            { return "focus" }
func (h *WindowFocusMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowFocusMethod) GetIsStatic() bool          { return true }
func (h *WindowFocusMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowFocusMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowFocusMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
