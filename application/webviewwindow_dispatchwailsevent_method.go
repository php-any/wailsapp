package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowDispatchWailsEventMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowDispatchWailsEventMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.DispatchWailsEvent 缺少参数, index: 0"))
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
			return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.DispatchWailsEvent 参数类型不支持, index: 0"))
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
			return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.DispatchWailsEvent 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(*applicationsrc.CustomEvent)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.DispatchWailsEvent 参数类型不支持, index: 0"))
	}

	h.source.DispatchWailsEvent(arg0)
	return nil, nil
}

func (h *WebviewWindowDispatchWailsEventMethod) GetName() string { return "dispatchWailsEvent" }
func (h *WebviewWindowDispatchWailsEventMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WebviewWindowDispatchWailsEventMethod) GetIsStatic() bool { return true }
func (h *WebviewWindowDispatchWailsEventMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "event", 0, nil, nil),
	}
}

func (h *WebviewWindowDispatchWailsEventMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "event", 0, nil),
	}
}

func (h *WebviewWindowDispatchWailsEventMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
