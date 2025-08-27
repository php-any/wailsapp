package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructCanSelectHiddenExtensionMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructCanSelectHiddenExtensionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.CanSelectHiddenExtension 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.CanSelectHiddenExtension(arg0)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *OpenFileDialogStructCanSelectHiddenExtensionMethod) GetName() string {
	return "canSelectHiddenExtension"
}
func (h *OpenFileDialogStructCanSelectHiddenExtensionMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogStructCanSelectHiddenExtensionMethod) GetIsStatic() bool { return true }
func (h *OpenFileDialogStructCanSelectHiddenExtensionMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "canSelectHiddenExtension", 0, nil, nil),
	}
}

func (h *OpenFileDialogStructCanSelectHiddenExtensionMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "canSelectHiddenExtension", 0, nil),
	}
}

func (h *OpenFileDialogStructCanSelectHiddenExtensionMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
