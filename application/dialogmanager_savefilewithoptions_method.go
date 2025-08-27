package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type DialogManagerSaveFileWithOptionsMethod struct {
	source *applicationsrc.DialogManager
}

func (h *DialogManagerSaveFileWithOptionsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("DialogManager.SaveFileWithOptions 缺少参数, index: 0"))
	}

	var arg0 *applicationsrc.SaveFileDialogOptions
	switch v := a0.(type) {
	case *data.ClassValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			// 检查 GetSource 返回的类型，如果是指针则解引用
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.SaveFileDialogOptions); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(*applicationsrc.SaveFileDialogOptions)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("DialogManager.SaveFileWithOptions 参数类型不支持, index: 0"))
		}
	case *data.ProxyValue:
		if p, ok := v.Class.(interface{ GetSource() any }); ok {
			if src := p.GetSource(); src != nil {
				if ptr, ok := src.(**applicationsrc.SaveFileDialogOptions); ok {
					arg0 = *ptr
				} else {
					arg0 = src.(*applicationsrc.SaveFileDialogOptions)
				}
			}
		} else {
			return nil, data.NewErrorThrow(nil, errors.New("DialogManager.SaveFileWithOptions 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(*applicationsrc.SaveFileDialogOptions)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("DialogManager.SaveFileWithOptions 参数类型不支持, index: 0"))
	}

	ret0 := h.source.SaveFileWithOptions(arg0)
	return data.NewClassValue(NewSaveFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *DialogManagerSaveFileWithOptionsMethod) GetName() string { return "saveFileWithOptions" }
func (h *DialogManagerSaveFileWithOptionsMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *DialogManagerSaveFileWithOptionsMethod) GetIsStatic() bool { return true }
func (h *DialogManagerSaveFileWithOptionsMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "options", 0, nil, nil),
	}
}

func (h *DialogManagerSaveFileWithOptionsMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "options", 0, nil),
	}
}

func (h *DialogManagerSaveFileWithOptionsMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
