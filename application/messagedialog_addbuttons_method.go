package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MessageDialogAddButtonsMethod struct {
	source *applicationsrc.MessageDialog
}

func (h *MessageDialogAddButtonsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MessageDialog.AddButtons 缺少参数, index: 0"))
	}

	arg0 := make([]*applicationsrc.Button, 0)
	for _, v := range a0.(*data.ArrayValue).Value {
		switch elemVal := v.(type) {
		case *data.ClassValue:
			if p, ok := elemVal.Class.(interface{ GetSource() any }); ok {
				// 检查 GetSource 返回的类型，如果是指针则解引用
				if src := p.GetSource(); src != nil {
					if ptr, ok := src.(**applicationsrc.Button); ok {
						arg0 = append(arg0, *ptr)
					} else {
						arg0 = append(arg0, src.(*applicationsrc.Button))
					}
				}
			} else {
				return nil, data.NewErrorThrow(nil, errors.New("MessageDialog.AddButtons 参数类型不支持, index: 0"))
			}
		case *data.ProxyValue:
			if p, ok := elemVal.Class.(interface{ GetSource() any }); ok {
				if src := p.GetSource(); src != nil {
					if ptr, ok := src.(**applicationsrc.Button); ok {
						arg0 = append(arg0, *ptr)
					} else {
						arg0 = append(arg0, src.(*applicationsrc.Button))
					}
				}
			} else {
				return nil, data.NewErrorThrow(nil, errors.New("MessageDialog.AddButtons 参数类型不支持, index: 0"))
			}
		case *data.AnyValue:
			arg0 = append(arg0, elemVal.Value.(*applicationsrc.Button))
		}
	}

	ret0 := h.source.AddButtons(arg0)
	return data.NewClassValue(NewMessageDialogClassFrom(ret0), ctx), nil
}

func (h *MessageDialogAddButtonsMethod) GetName() string            { return "addButtons" }
func (h *MessageDialogAddButtonsMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MessageDialogAddButtonsMethod) GetIsStatic() bool          { return true }
func (h *MessageDialogAddButtonsMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "buttons", 0, nil, nil),
	}
}

func (h *MessageDialogAddButtonsMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "buttons", 0, nil),
	}
}

func (h *MessageDialogAddButtonsMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
