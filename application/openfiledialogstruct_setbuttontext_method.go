package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructSetButtonTextMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructSetButtonTextMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.SetButtonText 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetButtonText(arg0)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *OpenFileDialogStructSetButtonTextMethod) GetName() string { return "setButtonText" }
func (h *OpenFileDialogStructSetButtonTextMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *OpenFileDialogStructSetButtonTextMethod) GetIsStatic() bool { return true }
func (h *OpenFileDialogStructSetButtonTextMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "text", 0, nil, nil),
	}
}

func (h *OpenFileDialogStructSetButtonTextMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "text", 0, nil),
	}
}

func (h *OpenFileDialogStructSetButtonTextMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
