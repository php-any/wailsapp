package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SaveFileDialogStructTreatsFilePackagesAsDirectoriesMethod struct {
	source *applicationsrc.SaveFileDialogStruct
}

func (h *SaveFileDialogStructTreatsFilePackagesAsDirectoriesMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.TreatsFilePackagesAsDirectories 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.TreatsFilePackagesAsDirectories(arg0)
	return data.NewClassValue(NewSaveFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *SaveFileDialogStructTreatsFilePackagesAsDirectoriesMethod) GetName() string {
	return "treatsFilePackagesAsDirectories"
}
func (h *SaveFileDialogStructTreatsFilePackagesAsDirectoriesMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *SaveFileDialogStructTreatsFilePackagesAsDirectoriesMethod) GetIsStatic() bool { return true }
func (h *SaveFileDialogStructTreatsFilePackagesAsDirectoriesMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "treatsFilePackagesAsDirectories", 0, nil, nil),
	}
}

func (h *SaveFileDialogStructTreatsFilePackagesAsDirectoriesMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "treatsFilePackagesAsDirectories", 0, nil),
	}
}

func (h *SaveFileDialogStructTreatsFilePackagesAsDirectoriesMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
