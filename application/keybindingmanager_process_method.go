package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type KeyBindingManagerProcessMethod struct {
	source *applicationsrc.KeyBindingManager
}

func (h *KeyBindingManagerProcessMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("KeyBindingManager.Process 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("KeyBindingManager.Process 缺少参数, index: 1"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	var arg1 applicationsrc.Window
	switch v := a1.(type) {
	case *data.ClassValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			// 检查 GetSource 返回的类型，如果是指针则解引用
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Window); ok {
					arg1 = *ptr
				} else {
					arg1 = src.(applicationsrc.Window)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("KeyBindingManager.Process 参数类型不支持, index: 1"))
		}
	case *data.ProxyValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Window); ok {
					arg1 = *ptr
				} else {
					arg1 = src.(applicationsrc.Window)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("KeyBindingManager.Process 参数类型不支持, index: 1"))
		}
	case *data.AnyValue:
		arg1 = v.Value.(applicationsrc.Window)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("KeyBindingManager.Process 参数类型不支持, index: 1"))
	}

	ret0 := h.source.Process(arg0, arg1)
	return data.NewBoolValue(ret0), nil
}

func (h *KeyBindingManagerProcessMethod) GetName() string            { return "process" }
func (h *KeyBindingManagerProcessMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *KeyBindingManagerProcessMethod) GetIsStatic() bool          { return true }
func (h *KeyBindingManagerProcessMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "accelerator", 0, nil, nil),
		node.NewParameter(nil, "window", 1, nil, nil),
	}
}

func (h *KeyBindingManagerProcessMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "accelerator", 0, nil),
		node.NewVariable(nil, "window", 1, nil),
	}
}

func (h *KeyBindingManagerProcessMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
