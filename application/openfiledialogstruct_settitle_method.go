package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructSetTitleMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructSetTitleMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.SetTitle 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	ret0 := h.source.SetTitle(arg0)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *OpenFileDialogStructSetTitleMethod) GetName() string            { return "setTitle" }
func (h *OpenFileDialogStructSetTitleMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *OpenFileDialogStructSetTitleMethod) GetIsStatic() bool          { return true }
func (h *OpenFileDialogStructSetTitleMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "title", 0, nil, nil),
	}
}

func (h *OpenFileDialogStructSetTitleMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "title", 0, nil),
	}
}

func (h *OpenFileDialogStructSetTitleMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
