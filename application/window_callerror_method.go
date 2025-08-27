package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowCallErrorMethod struct {
	source applicationsrc.Window
}

func (h *WindowCallErrorMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.CallError 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.CallError 缺少参数, index: 1"))
	}

	a2, ok := ctx.GetIndexValue(2)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.CallError 缺少参数, index: 2"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	arg1 := a1.(*data.StringValue).AsString()
	arg2, err := a2.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	h.source.CallError(arg0, arg1, arg2)
	return nil, nil
}

func (h *WindowCallErrorMethod) GetName() string            { return "callError" }
func (h *WindowCallErrorMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowCallErrorMethod) GetIsStatic() bool          { return true }
func (h *WindowCallErrorMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
		node.NewParameter(nil, "param1", 1, nil, nil),
		node.NewParameter(nil, "param2", 2, nil, nil),
	}
}

func (h *WindowCallErrorMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
		node.NewVariable(nil, "param1", 1, nil),
		node.NewVariable(nil, "param2", 2, nil),
	}
}

func (h *WindowCallErrorMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
