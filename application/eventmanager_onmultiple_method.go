package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type EventManagerOnMultipleMethod struct {
	source *applicationsrc.EventManager
}

func (h *EventManagerOnMultipleMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("EventManager.OnMultiple 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("EventManager.OnMultiple 缺少参数, index: 1"))
	}

	a2, ok := ctx.GetIndexValue(2)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("EventManager.OnMultiple 缺少参数, index: 2"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	var arg1 func(*applicationsrc.CustomEvent)
	switch fnv := a1.(type) {
	case *data.FuncValue:
		arg1 = func(p0 *applicationsrc.CustomEvent) { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("EventManager.OnMultiple 参数类型不支持, index: 1"))
	}
	arg2, err := a2.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	h.source.OnMultiple(arg0, arg1, arg2)
	return nil, nil
}

func (h *EventManagerOnMultipleMethod) GetName() string            { return "onMultiple" }
func (h *EventManagerOnMultipleMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *EventManagerOnMultipleMethod) GetIsStatic() bool          { return true }
func (h *EventManagerOnMultipleMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "name", 0, nil, nil),
		node.NewParameter(nil, "callback", 1, nil, nil),
		node.NewParameter(nil, "counter", 2, nil, nil),
	}
}

func (h *EventManagerOnMultipleMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "name", 0, nil),
		node.NewVariable(nil, "callback", 1, nil),
		node.NewVariable(nil, "counter", 2, nil),
	}
}

func (h *EventManagerOnMultipleMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
