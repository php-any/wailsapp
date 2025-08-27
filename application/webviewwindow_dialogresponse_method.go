package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowDialogResponseMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowDialogResponseMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.DialogResponse 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.DialogResponse 缺少参数, index: 1"))
	}

	a2, ok := ctx.GetIndexValue(2)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.DialogResponse 缺少参数, index: 2"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	arg1 := a1.(*data.StringValue).AsString()
	arg2, err := a2.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	h.source.DialogResponse(arg0, arg1, arg2)
	return nil, nil
}

func (h *WebviewWindowDialogResponseMethod) GetName() string            { return "dialogResponse" }
func (h *WebviewWindowDialogResponseMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowDialogResponseMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowDialogResponseMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "dialogID", 0, nil, nil),
		node.NewParameter(nil, "result", 1, nil, nil),
		node.NewParameter(nil, "isJSON", 2, nil, nil),
	}
}

func (h *WebviewWindowDialogResponseMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "dialogID", 0, nil),
		node.NewVariable(nil, "result", 1, nil),
		node.NewVariable(nil, "isJSON", 2, nil),
	}
}

func (h *WebviewWindowDialogResponseMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
