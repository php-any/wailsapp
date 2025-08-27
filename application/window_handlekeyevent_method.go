package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowHandleKeyEventMethod struct {
	source applicationsrc.Window
}

func (h *WindowHandleKeyEventMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.HandleKeyEvent 缺少参数, index: 0"))
	}

	arg0 := a0.(*data.StringValue).AsString()

	h.source.HandleKeyEvent(arg0)
	return nil, nil
}

func (h *WindowHandleKeyEventMethod) GetName() string            { return "handleKeyEvent" }
func (h *WindowHandleKeyEventMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowHandleKeyEventMethod) GetIsStatic() bool          { return true }
func (h *WindowHandleKeyEventMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *WindowHandleKeyEventMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *WindowHandleKeyEventMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
