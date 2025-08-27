package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowSetMaximiseButtonStateMethod struct {
	source applicationsrc.Window
}

func (h *WindowSetMaximiseButtonStateMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.SetMaximiseButtonState 缺少参数, index: 0"))
	}

	var arg0 applicationsrc.ButtonState
	arg0Int, err := a0.(*data.IntValue).AsInt()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}
	arg0 = applicationsrc.ButtonState(arg0Int)

	ret0 := h.source.SetMaximiseButtonState(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowSetMaximiseButtonStateMethod) GetName() string            { return "setMaximiseButtonState" }
func (h *WindowSetMaximiseButtonStateMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowSetMaximiseButtonStateMethod) GetIsStatic() bool          { return true }
func (h *WindowSetMaximiseButtonStateMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *WindowSetMaximiseButtonStateMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *WindowSetMaximiseButtonStateMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
