package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SaveFileDialogStructSetButtonTextMethod struct {
	source *applicationsrc.SaveFileDialogStruct
}

func (h *SaveFileDialogStructSetButtonTextMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.SetButtonText 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetButtonText(arg0)
	return data.NewClassValue(NewSaveFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *SaveFileDialogStructSetButtonTextMethod) GetName() string { return "setButtonText" }
func (h *SaveFileDialogStructSetButtonTextMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *SaveFileDialogStructSetButtonTextMethod) GetIsStatic() bool { return true }
func (h *SaveFileDialogStructSetButtonTextMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "text", 0, nil, nil),
	}
}

func (h *SaveFileDialogStructSetButtonTextMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "text", 0, nil),
	}
}

func (h *SaveFileDialogStructSetButtonTextMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
