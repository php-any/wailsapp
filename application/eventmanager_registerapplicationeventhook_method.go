package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

type EventManagerRegisterApplicationEventHookMethod struct {
	source *applicationsrc.EventManager
}

func (h *EventManagerRegisterApplicationEventHookMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("EventManager.RegisterApplicationEventHook 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("EventManager.RegisterApplicationEventHook 缺少参数, index: 1"))
	}

	var arg0 events.ApplicationEventType
	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 = events.ApplicationEventType(arg0Int)
	var arg1 func(*applicationsrc.ApplicationEvent)
	switch fnv := a1.(type) {
	case *data.FuncValue:
		arg1 = func(p0 *applicationsrc.ApplicationEvent) { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("EventManager.RegisterApplicationEventHook 参数类型不支持, index: 1"))
	}

	ret0 := h.source.RegisterApplicationEventHook(arg0, arg1)
	return data.NewAnyValue(ret0), nil
}

func (h *EventManagerRegisterApplicationEventHookMethod) GetName() string {
	return "registerApplicationEventHook"
}
func (h *EventManagerRegisterApplicationEventHookMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *EventManagerRegisterApplicationEventHookMethod) GetIsStatic() bool { return true }
func (h *EventManagerRegisterApplicationEventHookMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "eventType", 0, nil, nil),
		node.NewParameter(nil, "callback", 1, nil, nil),
	}
}

func (h *EventManagerRegisterApplicationEventHookMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "eventType", 0, nil),
		node.NewVariable(nil, "callback", 1, nil),
	}
}

func (h *EventManagerRegisterApplicationEventHookMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
