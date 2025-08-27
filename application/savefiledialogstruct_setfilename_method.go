package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SaveFileDialogStructSetFilenameMethod struct {
	source *applicationsrc.SaveFileDialogStruct
}

func (h *SaveFileDialogStructSetFilenameMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.SetFilename 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetFilename(arg0)
	return data.NewClassValue(NewSaveFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *SaveFileDialogStructSetFilenameMethod) GetName() string { return "setFilename" }
func (h *SaveFileDialogStructSetFilenameMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *SaveFileDialogStructSetFilenameMethod) GetIsStatic() bool { return true }
func (h *SaveFileDialogStructSetFilenameMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "filename", 0, nil, nil),
	}
}

func (h *SaveFileDialogStructSetFilenameMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "filename", 0, nil),
	}
}

func (h *SaveFileDialogStructSetFilenameMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
