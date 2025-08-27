package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowIsMaximisedMethod struct {
	source applicationsrc.Window
}

func (h *WindowIsMaximisedMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsMaximised()
	return data.NewBoolValue(ret0), nil
}

func (h *WindowIsMaximisedMethod) GetName() string            { return "isMaximised" }
func (h *WindowIsMaximisedMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowIsMaximisedMethod) GetIsStatic() bool          { return true }
func (h *WindowIsMaximisedMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *WindowIsMaximisedMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *WindowIsMaximisedMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
