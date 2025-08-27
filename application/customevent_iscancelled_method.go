package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type CustomEventIsCancelledMethod struct {
	source *applicationsrc.CustomEvent
}

func (h *CustomEventIsCancelledMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	ret0 := h.source.IsCancelled()
	return data.NewBoolValue(ret0), nil
}

func (h *CustomEventIsCancelledMethod) GetName() string            { return "isCancelled" }
func (h *CustomEventIsCancelledMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *CustomEventIsCancelledMethod) GetIsStatic() bool          { return true }
func (h *CustomEventIsCancelledMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *CustomEventIsCancelledMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *CustomEventIsCancelledMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
