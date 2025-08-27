package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuRemoveMenuItemMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuRemoveMenuItemMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Menu.RemoveMenuItem 缺少参数, index: 0"))
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
			return nil, data.NewErrorThrow(nil, errors.New("Menu.RemoveMenuItem 参数类型不支持, index: 0"))
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
			return nil, data.NewErrorThrow(nil, errors.New("Menu.RemoveMenuItem 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(*applicationsrc.MenuItem)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("Menu.RemoveMenuItem 参数类型不支持, index: 0"))
	}

	h.source.RemoveMenuItem(arg0)
	return nil, nil
}

func (h *MenuRemoveMenuItemMethod) GetName() string            { return "removeMenuItem" }
func (h *MenuRemoveMenuItemMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuRemoveMenuItemMethod) GetIsStatic() bool          { return true }
func (h *MenuRemoveMenuItemMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "target", 0, nil, nil),
	}
}

func (h *MenuRemoveMenuItemMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "target", 0, nil),
	}
}

func (h *MenuRemoveMenuItemMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
