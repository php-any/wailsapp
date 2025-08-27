package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowHandleWindowEventMethod struct {
	source applicationsrc.Window
}

func (h *WindowHandleWindowEventMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.HandleWindowEvent 缺少参数, index: 0"))
	}

	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 := uint(arg0Int)

	h.source.HandleWindowEvent(arg0)
	return nil, nil
}

func (h *WindowHandleWindowEventMethod) GetName() string            { return "handleWindowEvent" }
func (h *WindowHandleWindowEventMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowHandleWindowEventMethod) GetIsStatic() bool          { return true }
func (h *WindowHandleWindowEventMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *WindowHandleWindowEventMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *WindowHandleWindowEventMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
