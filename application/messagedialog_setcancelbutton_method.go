package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MessageDialogSetCancelButtonMethod struct {
	source *applicationsrc.MessageDialog
}

func (h *MessageDialogSetCancelButtonMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MessageDialog.SetCancelButton 缺少参数, index: 0"))
	}

	var arg0 *applicationsrc.Button
	switch v := a0.(type) {
	case *data.ClassValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			// 检查 GetSource 返回的类型，如果是指针则解引用
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.Button); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(*applicationsrc.Button)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("MessageDialog.SetCancelButton 参数类型不支持, index: 0"))
		}
	case *data.ProxyValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.Button); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(*applicationsrc.Button)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("MessageDialog.SetCancelButton 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(*applicationsrc.Button)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("MessageDialog.SetCancelButton 参数类型不支持, index: 0"))
	}

	ret0 := h.source.SetCancelButton(arg0)
	return data.NewClassValue(NewMessageDialogClassFrom(ret0), ctx), nil
}

func (h *MessageDialogSetCancelButtonMethod) GetName() string            { return "setCancelButton" }
func (h *MessageDialogSetCancelButtonMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MessageDialogSetCancelButtonMethod) GetIsStatic() bool          { return true }
func (h *MessageDialogSetCancelButtonMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "button", 0, nil, nil),
	}
}

func (h *MessageDialogSetCancelButtonMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "button", 0, nil),
	}
}

func (h *MessageDialogSetCancelButtonMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
