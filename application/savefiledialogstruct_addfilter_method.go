package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type SaveFileDialogStructAddFilterMethod struct {
	source *applicationsrc.SaveFileDialogStruct
}

func (h *SaveFileDialogStructAddFilterMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.AddFilter 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("SaveFileDialogStruct.AddFilter 缺少参数, index: 1"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	arg1 := a1.(*data.StringValue).AsString()

	ret0 := h.source.AddFilter(arg0, arg1)
	return data.NewClassValue(NewSaveFileDialogStructClassFrom(ret0), ctx), nil
}

func (h *SaveFileDialogStructAddFilterMethod) GetName() string            { return "addFilter" }
func (h *SaveFileDialogStructAddFilterMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *SaveFileDialogStructAddFilterMethod) GetIsStatic() bool          { return true }
func (h *SaveFileDialogStructAddFilterMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "displayName", 0, nil, nil),
		node.NewParameter(nil, "pattern", 1, nil, nil),
	}
}

func (h *SaveFileDialogStructAddFilterMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "displayName", 0, nil),
		node.NewVariable(nil, "pattern", 1, nil),
	}
}

func (h *SaveFileDialogStructAddFilterMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
