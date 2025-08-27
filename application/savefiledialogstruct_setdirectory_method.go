package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SaveFileDialogStructSetDirectoryMethod struct {
	source *applicationsrc.SaveFileDialogStruct
}

func (h *SaveFileDialogStructSetDirectoryMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.SetDirectory 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetDirectory(arg0)
	return data.NewClassValue(NewSaveFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *SaveFileDialogStructSetDirectoryMethod) GetName() string { return "setDirectory" }
func (h *SaveFileDialogStructSetDirectoryMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *SaveFileDialogStructSetDirectoryMethod) GetIsStatic() bool { return true }
func (h *SaveFileDialogStructSetDirectoryMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "directory", 0, nil, nil),
	}
}

func (h *SaveFileDialogStructSetDirectoryMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "directory", 0, nil),
	}
}

func (h *SaveFileDialogStructSetDirectoryMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
