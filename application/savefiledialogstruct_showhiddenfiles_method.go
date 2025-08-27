package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SaveFileDialogStructShowHiddenFilesMethod struct {
	source *applicationsrc.SaveFileDialogStruct
}

func (h *SaveFileDialogStructShowHiddenFilesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.ShowHiddenFiles 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.ShowHiddenFiles(arg0)
	return data.NewClassValue(NewSaveFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *SaveFileDialogStructShowHiddenFilesMethod) GetName() string { return "showHiddenFiles" }
func (h *SaveFileDialogStructShowHiddenFilesMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *SaveFileDialogStructShowHiddenFilesMethod) GetIsStatic() bool { return true }
func (h *SaveFileDialogStructShowHiddenFilesMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "showHiddenFiles", 0, nil, nil),
	}
}

func (h *SaveFileDialogStructShowHiddenFilesMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "showHiddenFiles", 0, nil),
	}
}

func (h *SaveFileDialogStructShowHiddenFilesMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
