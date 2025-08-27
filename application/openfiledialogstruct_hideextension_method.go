package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructHideExtensionMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructHideExtensionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.HideExtension 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.HideExtension(arg0)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *OpenFileDialogStructHideExtensionMethod) GetName() string { return "hideExtension" }
func (h *OpenFileDialogStructHideExtensionMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogStructHideExtensionMethod) GetIsStatic() bool { return true }
func (h *OpenFileDialogStructHideExtensionMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "hideExtension", 0, nil, nil),
	}
}

func (h *OpenFileDialogStructHideExtensionMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "hideExtension", 0, nil),
	}
}

func (h *OpenFileDialogStructHideExtensionMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
