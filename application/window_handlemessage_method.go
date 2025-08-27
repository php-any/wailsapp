package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowHandleMessageMethod struct {
	source applicationsrc.Window
}

func (h *WindowHandleMessageMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.HandleMessage 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	h.source.HandleMessage(arg0)
	return nil, nil
}

func (h *WindowHandleMessageMethod) GetName() string            { return "handleMessage" }
func (h *WindowHandleMessageMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowHandleMessageMethod) GetIsStatic() bool          { return true }
func (h *WindowHandleMessageMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *WindowHandleMessageMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *WindowHandleMessageMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
