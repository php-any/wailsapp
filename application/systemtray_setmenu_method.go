package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SystemTraySetMenuMethod struct {
	source *applicationsrc.SystemTray
}

func (h *SystemTraySetMenuMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.SetMenu 缺少参数, index: 0"))
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
			return nil, data.NewErrorThrow(nil, errors.New("SystemTray.SetMenu 参数类型不支持, index: 0"))
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
			return nil, data.NewErrorThrow(nil, errors.New("SystemTray.SetMenu 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(*applicationsrc.Menu)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("SystemTray.SetMenu 参数类型不支持, index: 0"))
	}

	ret0 := h.source.SetMenu(arg0)
	return data.NewClassValue(NewSystemTrayClassFrom(ret0), ctx), nil
}

func (h *SystemTraySetMenuMethod) GetName() string            { return "setMenu" }
func (h *SystemTraySetMenuMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SystemTraySetMenuMethod) GetIsStatic() bool          { return true }
func (h *SystemTraySetMenuMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "menu", 0, nil, nil),
	}
}

func (h *SystemTraySetMenuMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "menu", 0, nil),
	}
}

func (h *SystemTraySetMenuMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
