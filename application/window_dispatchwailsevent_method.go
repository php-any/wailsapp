package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowDispatchWailsEventMethod struct {
	source applicationsrc.Window
}

func (h *WindowDispatchWailsEventMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.DispatchWailsEvent 缺少参数, index: 0"))
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
			return nil, data.NewErrorThrow(nil, errors.New("Window.DispatchWailsEvent 参数类型不支持, index: 0"))
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
			return nil, data.NewErrorThrow(nil, errors.New("Window.DispatchWailsEvent 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(*applicationsrc.CustomEvent)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("Window.DispatchWailsEvent 参数类型不支持, index: 0"))
	}

	h.source.DispatchWailsEvent(arg0)
	return nil, nil
}

func (h *WindowDispatchWailsEventMethod) GetName() string            { return "dispatchWailsEvent" }
func (h *WindowDispatchWailsEventMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowDispatchWailsEventMethod) GetIsStatic() bool          { return true }
func (h *WindowDispatchWailsEventMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *WindowDispatchWailsEventMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *WindowDispatchWailsEventMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
