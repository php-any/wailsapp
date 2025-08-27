package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructCanCreateDirectoriesMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructCanCreateDirectoriesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.CanCreateDirectories 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.CanCreateDirectories(arg0)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *OpenFileDialogStructCanCreateDirectoriesMethod) GetName() string {
	return "canCreateDirectories"
}
func (h *OpenFileDialogStructCanCreateDirectoriesMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogStructCanCreateDirectoriesMethod) GetIsStatic() bool { return true }
func (h *OpenFileDialogStructCanCreateDirectoriesMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "canCreateDirectories", 0, nil, nil),
	}
}

func (h *OpenFileDialogStructCanCreateDirectoriesMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "canCreateDirectories", 0, nil),
	}
}

func (h *OpenFileDialogStructCanCreateDirectoriesMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
