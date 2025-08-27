package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuAddRadioMethod struct {
	source *applicationsrc.Menu
}

func (h *MenuAddRadioMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Menu.AddRadio 缺少参数, index: 0"))
	}

	a1, ok := ctx.GetIndexValue(1)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("Menu.AddRadio 缺少参数, index: 1"))
	}

	arg0 := a0.(*data.StringValue).AsString()
	arg1, err := a1.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.AddRadio(arg0, arg1)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *MenuAddRadioMethod) GetName() string            { return "addRadio" }
func (h *MenuAddRadioMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuAddRadioMethod) GetIsStatic() bool          { return true }
func (h *MenuAddRadioMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "label", 0, nil, nil),
		node.NewParameter(nil, "enabled", 1, nil, nil),
	}
}

func (h *MenuAddRadioMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "label", 0, nil),
		node.NewVariable(nil, "enabled", 1, nil),
	}
}

func (h *MenuAddRadioMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
