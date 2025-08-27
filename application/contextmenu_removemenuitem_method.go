package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type ContextMenuRemoveMenuItemMethod struct {
	source *applicationsrc.ContextMenu
}

func (h *ContextMenuRemoveMenuItemMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenu.RemoveMenuItem 缺少参数, index: 0"))
	}

	var arg0 *applicationsrc.MenuItem
	switch v := a0.(type) {
	case *data.ClassValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			// 检查 GetSource 返回的类型，如果是指针则解引用
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.MenuItem); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(*applicationsrc.MenuItem)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("ContextMenu.RemoveMenuItem 参数类型不支持, index: 0"))
		}
	case *data.ProxyValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.MenuItem); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(*applicationsrc.MenuItem)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("ContextMenu.RemoveMenuItem 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(*applicationsrc.MenuItem)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("ContextMenu.RemoveMenuItem 参数类型不支持, index: 0"))
	}

	h.source.RemoveMenuItem(arg0)
	return nil, nil
}

func (h *ContextMenuRemoveMenuItemMethod) GetName() string            { return "removeMenuItem" }
func (h *ContextMenuRemoveMenuItemMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *ContextMenuRemoveMenuItemMethod) GetIsStatic() bool          { return true }
func (h *ContextMenuRemoveMenuItemMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *ContextMenuRemoveMenuItemMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *ContextMenuRemoveMenuItemMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
