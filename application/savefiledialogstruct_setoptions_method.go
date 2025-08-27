package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SaveFileDialogStructSetOptionsMethod struct {
	source *applicationsrc.SaveFileDialogStruct
}

func (h *SaveFileDialogStructSetOptionsMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.SetOptions 缺少参数, index: 0"))
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
			return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.SetOptions 参数类型不支持, index: 0"))
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
			return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.SetOptions 参数类型不支持, index: 0"))
		}
	case *data.AnyValue:
		arg0 = v.Value.(*applicationsrc.SaveFileDialogOptions)
	default:
		return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.SetOptions 参数类型不支持, index: 0"))
	}

	h.source.SetOptions(arg0)
	return nil, nil
}

func (h *SaveFileDialogStructSetOptionsMethod) GetName() string { return "setOptions" }
func (h *SaveFileDialogStructSetOptionsMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *SaveFileDialogStructSetOptionsMethod) GetIsStatic() bool { return true }
func (h *SaveFileDialogStructSetOptionsMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "options", 0, nil, nil),
	}
}

func (h *SaveFileDialogStructSetOptionsMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "options", 0, nil),
	}
}

func (h *SaveFileDialogStructSetOptionsMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
