package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

type WindowOnWindowEventMethod struct {
	source applicationsrc.Window
}

func (h *WindowOnWindowEventMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.OnWindowEvent 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.OnWindowEvent 缺少参数, index: 1"))
	}

	var arg0 events.WindowEventType
	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 = events.WindowEventType(arg0Int)
	var arg1 func(*applicationsrc.WindowEvent)
	switch fnv := a1.(type) {
	case *data.FuncValue:
		arg1 = func(p0 *applicationsrc.WindowEvent) { fnv.Call(ctx) }
	default:
		return nil, data.NewErrorThrow(nil, errors.New("Window.OnWindowEvent 参数类型不支持, index: 1"))
	}

	ret0 := h.source.OnWindowEvent(arg0, arg1)
	return data.NewAnyValue(ret0), nil
}

func (h *WindowOnWindowEventMethod) GetName() string            { return "onWindowEvent" }
func (h *WindowOnWindowEventMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowOnWindowEventMethod) GetIsStatic() bool          { return true }
func (h *WindowOnWindowEventMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
		node.NewParameter(nil, "param1", 1, nil, nil),
	}
}

func (h *WindowOnWindowEventMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
		node.NewVariable(nil, "param1", 1, nil),
	}
}

func (h *WindowOnWindowEventMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
