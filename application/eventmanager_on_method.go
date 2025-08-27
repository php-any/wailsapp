package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type EventManagerOnMethod struct {
	source *applicationsrc.EventManager
}

func (h *EventManagerOnMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("EventManager.On 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("EventManager.On 缺少参数, index: 1"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	var arg1 func(*applicationsrc.CustomEvent)
	switch fnv := a1.(type) {
	case *data.FuncValue:
		arg1 = func(p0 *applicationsrc.CustomEvent) { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("EventManager.On 参数类型不支持, index: 1"))
	}

	ret0 := h.source.On(arg0, arg1)
	return data.NewAnyValue(ret0), nil
}

func (h *EventManagerOnMethod) GetName() string            { return "on" }
func (h *EventManagerOnMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *EventManagerOnMethod) GetIsStatic() bool          { return true }
func (h *EventManagerOnMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "name", 0, nil, nil),
		node.NewParameter(nil, "callback", 1, nil, nil),
	}
}

func (h *EventManagerOnMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "name", 0, nil),
		node.NewVariable(nil, "callback", 1, nil),
	}
}

func (h *EventManagerOnMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
