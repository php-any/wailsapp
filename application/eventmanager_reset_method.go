package application

import (
	"github.com/php-any/origami/data"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type EventManagerResetMethod struct {
	source *applicationsrc.EventManager
}

func (h *EventManagerResetMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	h.source.Reset()
	return nil, nil
}

func (h *EventManagerResetMethod) GetName() string            { return "reset" }
func (h *EventManagerResetMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *EventManagerResetMethod) GetIsStatic() bool          { return true }
func (h *EventManagerResetMethod) GetParams() []data.GetValue {
	return []data.GetValue{}
}

func (h *EventManagerResetMethod) GetVariables() []data.Variable {
	return []data.Variable{}
}

func (h *EventManagerResetMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
