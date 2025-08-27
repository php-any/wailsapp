package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructSetMessageMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructSetMessageMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.SetMessage 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetMessage(arg0)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *OpenFileDialogStructSetMessageMethod) GetName() string { return "setMessage" }
func (h *OpenFileDialogStructSetMessageMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogStructSetMessageMethod) GetIsStatic() bool { return true }
func (h *OpenFileDialogStructSetMessageMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "message", 0, nil, nil),
	}
}

func (h *OpenFileDialogStructSetMessageMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "message", 0, nil),
	}
}

func (h *OpenFileDialogStructSetMessageMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
