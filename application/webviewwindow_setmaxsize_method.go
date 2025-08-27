package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetMaxSizeMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetMaxSizeMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetMaxSize 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetMaxSize 缺少参数, index: 1"))
	}

	arg0, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg1, err := a1.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.SetMaxSize(arg0, arg1)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowSetMaxSizeMethod) GetName() string            { return "setMaxSize" }
func (h *WebviewWindowSetMaxSizeMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowSetMaxSizeMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowSetMaxSizeMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "maxWidth", 0, nil, nil),
		node.NewParameter(nil, "maxHeight", 1, nil, nil),
	}
}

func (h *WebviewWindowSetMaxSizeMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "maxWidth", 0, nil),
		node.NewVariable(nil, "maxHeight", 1, nil),
	}
}

func (h *WebviewWindowSetMaxSizeMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
