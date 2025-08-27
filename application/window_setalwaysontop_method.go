package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowSetAlwaysOnTopMethod struct {
	source applicationsrc.Window
}

func (h *WindowSetAlwaysOnTopMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.SetAlwaysOnTop 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.SetAlwaysOnTop(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowSetAlwaysOnTopMethod) GetName() string            { return "setAlwaysOnTop" }
func (h *WindowSetAlwaysOnTopMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowSetAlwaysOnTopMethod) GetIsStatic() bool          { return true }
func (h *WindowSetAlwaysOnTopMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *WindowSetAlwaysOnTopMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *WindowSetAlwaysOnTopMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
