package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SaveFileDialogStructSetMessageMethod struct {
	source *applicationsrc.SaveFileDialogStruct
}

func (h *SaveFileDialogStructSetMessageMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.SetMessage 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetMessage(arg0)
	return data.NewClassValue(NewSaveFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *SaveFileDialogStructSetMessageMethod) GetName() string { return "setMessage" }
func (h *SaveFileDialogStructSetMessageMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *SaveFileDialogStructSetMessageMethod) GetIsStatic() bool { return true }
func (h *SaveFileDialogStructSetMessageMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "message", 0, nil, nil),
	}
}

func (h *SaveFileDialogStructSetMessageMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "message", 0, nil),
	}
}

func (h *SaveFileDialogStructSetMessageMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
