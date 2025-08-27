package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type DialogManagerOpenFileWithOptionsMethod struct {
	source *applicationsrc.DialogManager
}

func (h *DialogManagerOpenFileWithOptionsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("DialogManager.OpenFileWithOptions 缺少参数, index: 0"))
	}

	var arg0 *applicationsrc.OpenFileDialogOptions
	switch v := a0.(type) {
	case *data.ClassValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			// 检查 GetSource 返回的类型，如果是指针则解引用
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.OpenFileDialogOptions); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(*applicationsrc.OpenFileDialogOptions)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("DialogManager.OpenFileWithOptions 参数类型不支持, index: 0"))
		}
	case *data.ProxyValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.OpenFileDialogOptions); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(*applicationsrc.OpenFileDialogOptions)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("DialogManager.OpenFileWithOptions 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(*applicationsrc.OpenFileDialogOptions)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("DialogManager.OpenFileWithOptions 参数类型不支持, index: 0"))
	}

	ret0 := h.source.OpenFileWithOptions(arg0)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *DialogManagerOpenFileWithOptionsMethod) GetName() string { return "openFileWithOptions" }
func (h *DialogManagerOpenFileWithOptionsMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *DialogManagerOpenFileWithOptionsMethod) GetIsStatic() bool { return true }
func (h *DialogManagerOpenFileWithOptionsMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "options", 0, nil, nil),
	}
}

func (h *DialogManagerOpenFileWithOptionsMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "options", 0, nil),
	}
}

func (h *DialogManagerOpenFileWithOptionsMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
