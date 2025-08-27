package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

type WindowRegisterHookMethod struct {
	source applicationsrc.Window
}

func (h *WindowRegisterHookMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.RegisterHook 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.RegisterHook 缺少参数, index: 1"))
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
		return nil, data.NewErrorThrow(nil, errors.New("Window.RegisterHook 参数类型不支持, index: 1"))
	}

	ret0 := h.source.RegisterHook(arg0, arg1)
	return data.NewAnyValue(ret0), nil
}

func (h *WindowRegisterHookMethod) GetName() string            { return "registerHook" }
func (h *WindowRegisterHookMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowRegisterHookMethod) GetIsStatic() bool          { return true }
func (h *WindowRegisterHookMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
		node.NewParameter(nil, "param1", 1, nil, nil),
	}
}

func (h *WindowRegisterHookMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
		node.NewVariable(nil, "param1", 1, nil),
	}
}

func (h *WindowRegisterHookMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
