package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructAttachToWindowMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructAttachToWindowMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.AttachToWindow 缺少参数, index: 0"))
	}

	var arg0 applicationsrc.Window
	switch v := a0.(type) {
	case *data.ClassValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			// 检查 GetSource 返回的类型，如果是指针则解引用
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Window); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(applicationsrc.Window)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.AttachToWindow 参数类型不支持, index: 0"))
		}
	case *data.ProxyValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(*applicationsrc.Window); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(applicationsrc.Window)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.AttachToWindow 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(applicationsrc.Window)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.AttachToWindow 参数类型不支持, index: 0"))
	}

	ret0 := h.source.AttachToWindow(arg0)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *OpenFileDialogStructAttachToWindowMethod) GetName() string { return "attachToWindow" }
func (h *OpenFileDialogStructAttachToWindowMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogStructAttachToWindowMethod) GetIsStatic() bool { return true }
func (h *OpenFileDialogStructAttachToWindowMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "window", 0, nil, nil),
	}
}

func (h *OpenFileDialogStructAttachToWindowMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "window", 0, nil),
	}
}

func (h *OpenFileDialogStructAttachToWindowMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
