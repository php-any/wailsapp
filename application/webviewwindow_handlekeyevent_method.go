package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowHandleKeyEventMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowHandleKeyEventMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.HandleKeyEvent 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	h.source.HandleKeyEvent(arg0)
	return nil, nil
}

func (h *WebviewWindowHandleKeyEventMethod) GetName() string            { return "handleKeyEvent" }
func (h *WebviewWindowHandleKeyEventMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowHandleKeyEventMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowHandleKeyEventMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "acceleratorString", 0, nil, nil),
	}
}

func (h *WebviewWindowHandleKeyEventMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "acceleratorString", 0, nil),
	}
}

func (h *WebviewWindowHandleKeyEventMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
