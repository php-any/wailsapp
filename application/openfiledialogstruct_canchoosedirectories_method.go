package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructCanChooseDirectoriesMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructCanChooseDirectoriesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.CanChooseDirectories 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.CanChooseDirectories(arg0)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *OpenFileDialogStructCanChooseDirectoriesMethod) GetName() string {
	return "canChooseDirectories"
}
func (h *OpenFileDialogStructCanChooseDirectoriesMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogStructCanChooseDirectoriesMethod) GetIsStatic() bool { return true }
func (h *OpenFileDialogStructCanChooseDirectoriesMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "canChooseDirectories", 0, nil, nil),
	}
}

func (h *OpenFileDialogStructCanChooseDirectoriesMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "canChooseDirectories", 0, nil),
	}
}

func (h *OpenFileDialogStructCanChooseDirectoriesMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
