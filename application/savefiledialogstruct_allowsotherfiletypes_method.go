package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SaveFileDialogStructAllowsOtherFileTypesMethod struct {
	source *applicationsrc.SaveFileDialogStruct
}

func (h *SaveFileDialogStructAllowsOtherFileTypesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.AllowsOtherFileTypes 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.AllowsOtherFileTypes(arg0)
	return data.NewClassValue(NewSaveFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *SaveFileDialogStructAllowsOtherFileTypesMethod) GetName() string {
	return "allowsOtherFileTypes"
}
func (h *SaveFileDialogStructAllowsOtherFileTypesMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *SaveFileDialogStructAllowsOtherFileTypesMethod) GetIsStatic() bool { return true }
func (h *SaveFileDialogStructAllowsOtherFileTypesMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "allowOtherFileTypes", 0, nil, nil),
	}
}

func (h *SaveFileDialogStructAllowsOtherFileTypesMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "allowOtherFileTypes", 0, nil),
	}
}

func (h *SaveFileDialogStructAllowsOtherFileTypesMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
