package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WebviewWindowHandleMessageMethod struct {
	source *applicationsrc.WebviewWindow
}

func (h *WebviewWindowHandleMessageMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("WebviewWindow.HandleMessage 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	h.source.HandleMessage(arg0)
	return nil, nil
}

func (h *WebviewWindowHandleMessageMethod) GetName() string            { return "handleMessage" }
func (h *WebviewWindowHandleMessageMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WebviewWindowHandleMessageMethod) GetIsStatic() bool          { return true }
func (h *WebviewWindowHandleMessageMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "message", 0, nil, nil),
	}
}

func (h *WebviewWindowHandleMessageMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "message", 0, nil),
	}
}

func (h *WebviewWindowHandleMessageMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
