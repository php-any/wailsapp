package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowSetMinSizeMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowSetMinSizeMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetMinSize 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.SetMinSize 缺少参数, index: 1"))
	}

	arg0, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg1, err := a1.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.SetMinSize(arg0, arg1)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WebviewWindowSetMinSizeMethod) GetName() string            { return "setMinSize" }
func (h *WebviewWindowSetMinSizeMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowSetMinSizeMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowSetMinSizeMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "minWidth", 0, nil, nil),
		node.NewParameter(nil, "minHeight", 1, nil, nil),
	}
}

func (h *WebviewWindowSetMinSizeMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "minWidth", 0, nil),
		node.NewVariable(nil, "minHeight", 1, nil),
	}
}

func (h *WebviewWindowSetMinSizeMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
