package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowInitiateFrontendDropProcessingMethod struct {
	source applicationsrc.Window
}

func (h *WindowInitiateFrontendDropProcessingMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.InitiateFrontendDropProcessing 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.InitiateFrontendDropProcessing 缺少参数, index: 1"))
	}

	a2, ok := ctx.GetIndexValue(2)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.InitiateFrontendDropProcessing 缺少参数, index: 2"))
	}

	arg0 := make([]string, 0)
	for _, v := range a0.(*data.ArrayValue).Value {
		switch elemVal := v.(type) {
		case *data.ClassValue:
			if p, ok := elemVal.Class.(interface{ GetSource() any }); ok {
				// 检查 GetSource 返回的类型，如果是指针则解引用
				if src := p.GetSource(); src != nil {
					if ptr, ok := src.(*string); ok {
						arg0 = append(arg0, *ptr)
					} else {
						arg0 = append(arg0, src.(string))
					}
				}
			} else {
				return nil, data.NewErrorThrow(nil, errors.New("Window.InitiateFrontendDropProcessing 参数类型不支持, index: 0"))
			}
		case *data.ProxyValue:
			if p, ok := elemVal.Class.(interface{ GetSource() any }); ok {
				if src := p.GetSource(); src != nil {
					if ptr, ok := src.(*string); ok {
						arg0 = append(arg0, *ptr)
					} else {
						arg0 = append(arg0, src.(string))
					}
				}
			} else {
				return nil, data.NewErrorThrow(nil, errors.New("Window.InitiateFrontendDropProcessing 参数类型不支持, index: 0"))
			}
		case *data.AnyValue:
			arg0 = append(arg0, elemVal.Value.(string))
		}
	}
	arg1, err := a1.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg2, err := a2.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	h.source.InitiateFrontendDropProcessing(arg0, arg1, arg2)
	return nil, nil
}

func (h *WindowInitiateFrontendDropProcessingMethod) GetName() string {
	return "initiateFrontendDropProcessing"
}
func (h *WindowInitiateFrontendDropProcessingMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WindowInitiateFrontendDropProcessingMethod) GetIsStatic() bool { return true }
func (h *WindowInitiateFrontendDropProcessingMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
		node.NewParameter(nil, "param1", 1, nil, nil),
		node.NewParameter(nil, "param2", 2, nil, nil),
	}
}

func (h *WindowInitiateFrontendDropProcessingMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
		node.NewVariable(nil, "param1", 1, nil),
		node.NewVariable(nil, "param2", 2, nil),
	}
}

func (h *WindowInitiateFrontendDropProcessingMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
