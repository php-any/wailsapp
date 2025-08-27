package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructTreatsFilePackagesAsDirectoriesMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructTreatsFilePackagesAsDirectoriesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.TreatsFilePackagesAsDirectories 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.TreatsFilePackagesAsDirectories(arg0)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *OpenFileDialogStructTreatsFilePackagesAsDirectoriesMethod) GetName() string {
	return "treatsFilePackagesAsDirectories"
}
func (h *OpenFileDialogStructTreatsFilePackagesAsDirectoriesMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogStructTreatsFilePackagesAsDirectoriesMethod) GetIsStatic() bool { return true }
func (h *OpenFileDialogStructTreatsFilePackagesAsDirectoriesMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "treatsFilePackagesAsDirectories", 0, nil, nil),
	}
}

func (h *OpenFileDialogStructTreatsFilePackagesAsDirectoriesMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "treatsFilePackagesAsDirectories", 0, nil),
	}
}

func (h *OpenFileDialogStructTreatsFilePackagesAsDirectoriesMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
