package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MessageDialogSetIconMethod struct {
	source *applicationsrc.MessageDialog
}

func (h *MessageDialogSetIconMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MessageDialog.SetIcon 缺少参数, index: 0"))
	}

	arg0 := make([]uint8, 0)
	for _, v := range a0.(*data.ArrayValue).Value {
		switch elemVal := v.(type) {
		case *data.ClassValue:
			if p, ok := elemVal.Class.(interface{ GetSource() any }); ok {
				// 检查 GetSource 返回的类型，如果是指针则解引用
				if src := p.GetSource(); src != nil {
					if ptr, ok := src.(*uint8); ok {
						arg0 = append(arg0, *ptr)
					} else {
						arg0 = append(arg0, src.(uint8))
					}
				}
			} else {
				return nil, data.NewErrorThrow(nil, errors.New("MessageDialog.SetIcon 参数类型不支持, index: 0"))
			}
		case *data.ProxyValue:
			if p, ok := elemVal.Class.(interface{ GetSource() any }); ok {
				if src := p.GetSource(); src != nil {
					if ptr, ok := src.(*uint8); ok {
						arg0 = append(arg0, *ptr)
					} else {
						arg0 = append(arg0, src.(uint8))
					}
				}
			} else {
				return nil, data.NewErrorThrow(nil, errors.New("MessageDialog.SetIcon 参数类型不支持, index: 0"))
			}
		case *data.AnyValue:
			arg0 = append(arg0, elemVal.Value.(uint8))
		}
	}

	ret0 := h.source.SetIcon(arg0)
	return data.NewClassValue(NewMessageDialogClassFrom(ret0), ctx), nil
}

func (h *MessageDialogSetIconMethod) GetName() string            { return "setIcon" }
func (h *MessageDialogSetIconMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MessageDialogSetIconMethod) GetIsStatic() bool          { return true }
func (h *MessageDialogSetIconMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "icon", 0, nil, nil),
	}
}

func (h *MessageDialogSetIconMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "icon", 0, nil),
	}
}

func (h *MessageDialogSetIconMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
