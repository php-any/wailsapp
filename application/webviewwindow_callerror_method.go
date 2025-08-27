package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowCallErrorMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowCallErrorMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.CallError 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.CallError 缺少参数, index: 1"))
	}

	a2, ok := ctx.GetIndexValue(2)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.CallError 缺少参数, index: 2"))
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

func (h *WebviewWindowCallErrorMethod) GetName() string            { return "callError" }
func (h *WebviewWindowCallErrorMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowCallErrorMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowCallErrorMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "callID", 0, nil, nil),
		node.NewParameter(nil, "result", 1, nil, nil),
		node.NewParameter(nil, "isJSON", 2, nil, nil),
	}
}

func (h *WebviewWindowCallErrorMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "callID", 0, nil),
		node.NewVariable(nil, "result", 1, nil),
		node.NewVariable(nil, "isJSON", 2, nil),
	}
}

func (h *WebviewWindowCallErrorMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
