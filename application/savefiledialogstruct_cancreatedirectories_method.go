package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SaveFileDialogStructCanCreateDirectoriesMethod struct {
	source *applicationsrc.SaveFileDialogStruct
}

func (h *SaveFileDialogStructCanCreateDirectoriesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.CanCreateDirectories 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.CanCreateDirectories(arg0)
	return data.NewClassValue(NewSaveFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *SaveFileDialogStructCanCreateDirectoriesMethod) GetName() string {
	return "canCreateDirectories"
}
func (h *SaveFileDialogStructCanCreateDirectoriesMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *SaveFileDialogStructCanCreateDirectoriesMethod) GetIsStatic() bool { return true }
func (h *SaveFileDialogStructCanCreateDirectoriesMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "canCreateDirectories", 0, nil, nil),
	}
}

func (h *SaveFileDialogStructCanCreateDirectoriesMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "canCreateDirectories", 0, nil),
	}
}

func (h *SaveFileDialogStructCanCreateDirectoriesMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
