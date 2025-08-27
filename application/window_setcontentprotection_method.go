package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type WindowSetContentProtectionMethod struct {
	source applicationsrc.Window
}

func (h *WindowSetContentProtectionMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Window.SetContentProtection 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.SetContentProtection(arg0)
	return data.NewClassValue(NewWindowClassFrom(ret0), ctx), nil
}

func (h *WindowSetContentProtectionMethod) GetName() string            { return "setContentProtection" }
func (h *WindowSetContentProtectionMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *WindowSetContentProtectionMethod) GetIsStatic() bool          { return true }
func (h *WindowSetContentProtectionMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "param0", 0, nil, nil),
	}
}

func (h *WindowSetContentProtectionMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "param0", 0, nil),
	}
}

func (h *WindowSetContentProtectionMethod) GetReturnType() data.Types {
	return data.NewBaseType("void")
}
