package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowIsMinimisedMethod struct {
	source applicationsrc.Window
}

func (h *WindowIsMinimisedMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsMinimised()
	return data.NewBoolValue(ret0), nil
}

func (h *WindowIsMinimisedMethod) GetName() string            { return "isMinimised" }
func (h *WindowIsMinimisedMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowIsMinimisedMethod) GetIsStatic() bool          { return true }
func (h *WindowIsMinimisedMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowIsMinimisedMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowIsMinimisedMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
