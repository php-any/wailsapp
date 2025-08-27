package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuManagerSetMethod struct {
	source *applicationsrc.MenuManager
}

func (h *MenuManagerSetMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MenuManager.Set 缺少参数, index: 0"))
	}

	var arg0 *applicationsrc.Menu
	switch v := a0.(type) {
	case *data.ClassValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			// 检查 GetSource 返回的类型，如果是指针则解引用
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.Menu); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(*applicationsrc.Menu)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("MenuManager.Set 参数类型不支持, index: 0"))
		}
	case *data.ProxyValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.Menu); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(*applicationsrc.Menu)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("MenuManager.Set 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(*applicationsrc.Menu)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("MenuManager.Set 参数类型不支持, index: 0"))
	}

	h.source.Set(arg0)
	return nil, nil
}

func (h *MenuManagerSetMethod) GetName() string            { return "set" }
func (h *MenuManagerSetMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuManagerSetMethod) GetIsStatic() bool          { return true }
func (h *MenuManagerSetMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "menu", 0, nil, nil),
	}
}

func (h *MenuManagerSetMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "menu", 0, nil),
	}
}

func (h *MenuManagerSetMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
