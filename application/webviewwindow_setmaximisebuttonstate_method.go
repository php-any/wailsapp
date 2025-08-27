package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetMaximiseButtonStateMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetMaximiseButtonStateMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetMaximiseButtonState 缺少参数, index: 0"))
	}

	var arg0 applicationsrc.ButtonState
	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 = applicationsrc.ButtonState(arg0Int)

	ret0 := h.source.SetMaximiseButtonState(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowSetMaximiseButtonStateMethod) GetName() string { return "setMaximiseButtonState" }
func (h *WebviewWindowSetMaximiseButtonStateMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WebviewWindowSetMaximiseButtonStateMethod) GetIsStatic() bool { return true }
func (h *WebviewWindowSetMaximiseButtonStateMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "state", 0, nil, nil),
	}
}

func (h *WebviewWindowSetMaximiseButtonStateMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "state", 0, nil),
	}
}

func (h *WebviewWindowSetMaximiseButtonStateMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
