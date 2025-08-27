package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetRelativePositionMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetRelativePositionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetRelativePosition 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetRelativePosition 缺少参数, index: 1"))
	}

	arg0, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg1, err := a1.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.SetRelativePosition(arg0, arg1)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowSetRelativePositionMethod) GetName() string { return "setRelativePosition" }
func (h *WebviewWindowSetRelativePositionMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WebviewWindowSetRelativePositionMethod) GetIsStatic() bool { return true }
func (h *WebviewWindowSetRelativePositionMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "x", 0, nil, nil),
		node.NewParameter(nil, "y", 1, nil, nil),
	}
}

func (h *WebviewWindowSetRelativePositionMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "x", 0, nil),
		node.NewVariable(nil, "y", 1, nil),
	}
}

func (h *WebviewWindowSetRelativePositionMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
