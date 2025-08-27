package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowDialogErrorMethod struct {
	source applicationsrc.Window
}

func (h *WindowDialogErrorMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.DialogError 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.DialogError 缺少参数, index: 1"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	arg1 := a1.(*data.StringValue).AsString()

	h.source.DialogError(arg0, arg1)
	return nil, nil
}

func (h *WindowDialogErrorMethod) GetName() string            { return "dialogError" }
func (h *WindowDialogErrorMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowDialogErrorMethod) GetIsStatic() bool          { return true }
func (h *WindowDialogErrorMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
		node.NewParameter(nil, "param1", 1, nil, nil),
	}
}

func (h *WindowDialogErrorMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
		node.NewVariable(nil, "param1", 1, nil),
	}
}

func (h *WindowDialogErrorMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
