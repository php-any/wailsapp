package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetSizeMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetSizeMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetSize 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetSize 缺少参数, index: 1"))
	}

	arg0, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg1, err := a1.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.SetSize(arg0, arg1)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowSetSizeMethod) GetName() string            { return "setSize" }
func (h *WebviewWindowSetSizeMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowSetSizeMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowSetSizeMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "width", 0, nil, nil),
		node.NewParameter(nil, "height", 1, nil, nil),
	}
}

func (h *WebviewWindowSetSizeMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "width", 0, nil),
		node.NewVariable(nil, "height", 1, nil),
	}
}

func (h *WebviewWindowSetSizeMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
