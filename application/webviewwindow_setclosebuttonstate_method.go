package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetCloseButtonStateMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetCloseButtonStateMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetCloseButtonState 缺少参数, index: 0"))
	}

	var arg0 applicationsrc.ButtonState
	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 = applicationsrc.ButtonState(arg0Int)

	ret0 := h.source.SetCloseButtonState(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowSetCloseButtonStateMethod) GetName() string { return "setCloseButtonState" }
func (h *WebviewWindowSetCloseButtonStateMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WebviewWindowSetCloseButtonStateMethod) GetIsStatic() bool { return true }
func (h *WebviewWindowSetCloseButtonStateMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "state", 0, nil, nil),
	}
}

func (h *WebviewWindowSetCloseButtonStateMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "state", 0, nil),
	}
}

func (h *WebviewWindowSetCloseButtonStateMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
