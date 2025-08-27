package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SaveFileDialogStructHideExtensionMethod struct {
	source *applicationsrc.SaveFileDialogStruct
}

func (h *SaveFileDialogStructHideExtensionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.HideExtension 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.HideExtension(arg0)
	return data.NewClassValue(NewSaveFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *SaveFileDialogStructHideExtensionMethod) GetName() string { return "hideExtension" }
func (h *SaveFileDialogStructHideExtensionMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *SaveFileDialogStructHideExtensionMethod) GetIsStatic() bool { return true }
func (h *SaveFileDialogStructHideExtensionMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "hideExtension", 0, nil, nil),
	}
}

func (h *SaveFileDialogStructHideExtensionMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "hideExtension", 0, nil),
	}
}

func (h *SaveFileDialogStructHideExtensionMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
