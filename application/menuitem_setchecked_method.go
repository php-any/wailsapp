package application

import (
	"errors"
	"github.com/php-any/origami/data"
	"github.com/php-any/origami/node"
	applicationsrc "github.com/wailsapp/wails/v3/pkg/application"
)

type MenuItemSetCheckedMethod struct {
	source *applicationsrc.MenuItem
}

func (h *MenuItemSetCheckedMethod) Call(ctx data.Context) (data.GetValue, data.Control) {

	a0, ok := ctx.GetIndexValue(0)
	if !ok {
		return nil, data.NewErrorThrow(nil, errors.New("MenuItem.SetChecked 缺少参数, index: 0"))
	}

	arg0, err := a0.(*data.BoolValue).AsBool()
	if err != nil {
		return nil, data.NewErrorThrow(nil, err)
	}

	ret0 := h.source.SetChecked(arg0)
	return data.NewClassValue(NewMenuItemClassFrom(ret0), ctx), nil
}

func (h *MenuItemSetCheckedMethod) GetName() string            { return "setChecked" }
func (h *MenuItemSetCheckedMethod) GetModifier() data.Modifier { return data.ModifierPublic }
func (h *MenuItemSetCheckedMethod) GetIsStatic() bool          { return true }
func (h *MenuItemSetCheckedMethod) GetParams() []data.GetValue {
	return []data.GetValue{
		node.NewParameter(nil, "checked", 0, nil, nil),
	}
}

func (h *MenuItemSetCheckedMethod) GetVariables() []data.Variable {
	return []data.Variable{
		node.NewVariable(nil, "checked", 0, nil),
	}
}

func (h *MenuItemSetCheckedMethod) GetReturnType() data.Types { return data.NewBaseType("void") }
