package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowErrorMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowErrorMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.Error 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.Error 缺少参数, index: 1"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	arg1 := make([]any, 0)
	for _, v := range a1.(*data.ArrayValue).Value {
		switch elemVal := v.(type) {
		case *data.ClassValue:
			if p, ok := elemVal.Class.(interface{ GetSource() any }); ok {
				// 检查 GetSource 返回的类型，如果是指针则解引用
				if src := p.GetSource(); src != nil {
					if ptr, ok := src.(*any); ok {
						arg1 = append(arg1, *ptr)
					} else {
						arg1 = append(arg1, src.(any))
					}
				}
			} else {
				return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.Error 参数类型不支持, index: 1"))
			}
		case *data.ProxyValue:
			if p, ok := elemVal.Class.(interface{ GetSource() any }); ok {
				if src := p.GetSource(); src != nil {
					if ptr, ok := src.(*any); ok {
						arg1 = append(arg1, *ptr)
					} else {
						arg1 = append(arg1, src.(any))
					}
				}
			} else {
				return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.Error 参数类型不支持, index: 1"))
			}
		case *data.AnyValue:
			arg1 = append(arg1, elemVal.Value.(any))
		}
	}

	h.source.Error(arg0, arg1...)
	return nil, nil
}

func (h *WebviewWindowErrorMethod) GetName() string            { return "error" }
func (h *WebviewWindowErrorMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowErrorMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowErrorMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "message", 0, nil, nil),
		node.NewParameters(nil, "args", 1, nil, nil),
	}
}

func (h *WebviewWindowErrorMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "message", 0, nil),
		node.NewVariable(nil, "args", 1, nil),
	}
}

func (h *WebviewWindowErrorMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
