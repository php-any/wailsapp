package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructCanChooseFilesMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructCanChooseFilesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.CanChooseFiles 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.CanChooseFiles(arg0)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *OpenFileDialogStructCanChooseFilesMethod) GetName() string { return "canChooseFiles" }
func (h *OpenFileDialogStructCanChooseFilesMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogStructCanChooseFilesMethod) GetIsStatic() bool { return true }
func (h *OpenFileDialogStructCanChooseFilesMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "canChooseFiles", 0, nil, nil),
	}
}

func (h *OpenFileDialogStructCanChooseFilesMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "canChooseFiles", 0, nil),
	}
}

func (h *OpenFileDialogStructCanChooseFilesMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
