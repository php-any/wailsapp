package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructShowHiddenFilesMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructShowHiddenFilesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.ShowHiddenFiles 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.ShowHiddenFiles(arg0)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *OpenFileDialogStructShowHiddenFilesMethod) GetName() string { return "showHiddenFiles" }
func (h *OpenFileDialogStructShowHiddenFilesMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogStructShowHiddenFilesMethod) GetIsStatic() bool { return true }
func (h *OpenFileDialogStructShowHiddenFilesMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "showHiddenFiles", 0, nil, nil),
	}
}

func (h *OpenFileDialogStructShowHiddenFilesMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "showHiddenFiles", 0, nil),
	}
}

func (h *OpenFileDialogStructShowHiddenFilesMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
