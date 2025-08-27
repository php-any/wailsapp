package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type OpenFileDialogStructAddFilterMethod struct {
	source *applicationsrc.OpenFileDialogStruct
}

func (h *OpenFileDialogStructAddFilterMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.AddFilter 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("OpenFileDialogStruct.AddFilter 缺少参数, index: 1"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	arg1 := a1.(*data.StringValue).AsString()

	ret0 := h.source.AddFilter(arg0, arg1)
	return data.NewClassValue(NewOpenFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *OpenFileDialogStructAddFilterMethod) GetName() string            { return "addFilter" }
func (h *OpenFileDialogStructAddFilterMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *OpenFileDialogStructAddFilterMethod) GetIsStatic() bool          { return true }
func (h *OpenFileDialogStructAddFilterMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "displayName", 0, nil, nil),
		node.NewParameter(nil, "pattern", 1, nil, nil),
	}
}

func (h *OpenFileDialogStructAddFilterMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "displayName", 0, nil),
		node.NewVariable(nil, "pattern", 1, nil),
	}
}

func (h *OpenFileDialogStructAddFilterMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
