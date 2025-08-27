package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type EventManagerOffMethod struct {
	source *applicationsrc.EventManager
}

func (h *EventManagerOffMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("EventManager.Off 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	h.source.Off(arg0)
	return nil, nil
}

func (h *EventManagerOffMethod) GetName() string            { return "off" }
func (h *EventManagerOffMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *EventManagerOffMethod) GetIsStatic() bool          { return true }
func (h *EventManagerOffMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "name", 0, nil, nil),
	}
}

func (h *EventManagerOffMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "name", 0, nil),
	}
}

func (h *EventManagerOffMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
