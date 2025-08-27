package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type EventManagerEmitEventMethod struct {
	source *applicationsrc.EventManager
}

func (h *EventManagerEmitEventMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("EventManager.EmitEvent 缺少参数, index: 0"))
	}

	var arg0 *applicationsrc.CustomEvent
	switch v := a0.(type) {
	case *data.ClassValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			// 检查 GetSource 返回的类型，如果是指针则解引用
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.CustomEvent); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(*applicationsrc.CustomEvent)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("EventManager.EmitEvent 参数类型不支持, index: 0"))
		}
	case *data.ProxyValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.CustomEvent); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(*applicationsrc.CustomEvent)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("EventManager.EmitEvent 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(*applicationsrc.CustomEvent)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("EventManager.EmitEvent 参数类型不支持, index: 0"))
	}

	h.source.EmitEvent(arg0)
	return nil, nil
}

func (h *EventManagerEmitEventMethod) GetName() string            { return "emitEvent" }
func (h *EventManagerEmitEventMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *EventManagerEmitEventMethod) GetIsStatic() bool          { return true }
func (h *EventManagerEmitEventMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "event", 0, nil, nil),
	}
}

func (h *EventManagerEmitEventMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "event", 0, nil),
	}
}

func (h *EventManagerEmitEventMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
