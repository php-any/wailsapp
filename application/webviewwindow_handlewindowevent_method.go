package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowHandleWindowEventMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowHandleWindowEventMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.HandleWindowEvent 缺少参数, index: 0"))
	}

	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 := uint(arg0Int)

	h.source.HandleWindowEvent(arg0)
	return nil, nil
}

func (h *WebviewWindowHandleWindowEventMethod) GetName() string { return "handleWindowEvent" }
func (h *WebviewWindowHandleWindowEventMethod) GetModifier() data.Modifier {
	return data.ModifierPublic
}
func (h *WebviewWindowHandleWindowEventMethod) GetIsStatic() bool { return true }
func (h *WebviewWindowHandleWindowEventMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "id", 0, nil, nil),
	}
}

func (h *WebviewWindowHandleWindowEventMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "id", 0, nil),
	}
}

func (h *WebviewWindowHandleWindowEventMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
