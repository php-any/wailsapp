package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuManagerAddMethod struct {
	source *applicationsrc.ContextMenuManager
}

func (h *ContextMenuManagerAddMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenuManager.Add 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenuManager.Add 缺少参数, index: 1"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	var arg1 *applicationsrc.ContextMenu
	switch v := a1.(type) {
	case *data.ClassValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			// 检查 GetSource 返回的类型，如果是指针则解引用
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.ContextMenu); ok {
					arg1 = *ptr
				} else {
					arg1 = src.(*applicationsrc.ContextMenu)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("ContextMenuManager.Add 参数类型不支持, index: 1"))
		}
	case *data.ProxyValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.ContextMenu); ok {
					arg1 = *ptr
				} else {
					arg1 = src.(*applicationsrc.ContextMenu)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("ContextMenuManager.Add 参数类型不支持, index: 1"))
		}
	case *data.AnyValue:
		arg1 = v.Value.(*applicationsrc.ContextMenu)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenuManager.Add 参数类型不支持, index: 1"))
	}

	h.source.Add(arg0, arg1)
	return nil, nil
}

func (h *ContextMenuManagerAddMethod) GetName() string            { return "add" }
func (h *ContextMenuManagerAddMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuManagerAddMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuManagerAddMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "name", 0, nil, nil),
		node.NewParameter(nil, "menu", 1, nil, nil),
	}
}

func (h *ContextMenuManagerAddMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "name", 0, nil),
		node.NewVariable(nil, "menu", 1, nil),
	}
}

func (h *ContextMenuManagerAddMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
