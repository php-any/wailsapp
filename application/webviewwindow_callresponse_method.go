package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowCallResponseMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowCallResponseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.CallResponse 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.CallResponse 缺少参数, index: 1"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	arg1 := a1.(*data.StringValue).AsString()

	h.source.CallResponse(arg0, arg1)
	return nil, nil
}

func (h *WebviewWindowCallResponseMethod) GetName() string            { return "callResponse" }
func (h *WebviewWindowCallResponseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowCallResponseMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowCallResponseMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "callID", 0, nil, nil),
		node.NewParameter(nil, "result", 1, nil, nil),
	}
}

func (h *WebviewWindowCallResponseMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "callID", 0, nil),
		node.NewVariable(nil, "result", 1, nil),
	}
}

func (h *WebviewWindowCallResponseMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
